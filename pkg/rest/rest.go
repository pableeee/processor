package rest

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pableeee/processor/pkg/limiter"
)

var (
	ErrNotFound        = errors.New("user not found")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnknownError    = errors.New("unknown error")
)

type Adapter[T any] interface {
	ParseResponse([]byte) (*T, error)
}

type GenericAdapter[T any] int

func (g *GenericAdapter[T]) ParseResponse(b []byte) (*T, error) {
	r := new(T)
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, nil
}

type Client[T any] interface {
	Execute(method, url string, body io.Reader, headers map[string]string) (*T, error)
	ExecuteWithContext(ctx context.Context, method, url string, body io.Reader, headers map[string]string) (*T, error)
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client[T any] struct {
	ratelimiter limiter.RateLimiter
	client      httpClient
	adapter     Adapter[T]
}

func NewRestClient[T any]() *client[T] {
	v := &client[T]{adapter: new(GenericAdapter[T])}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	v.client = &http.Client{Transport: tr}

	return v
}
func WithLimiter[T any](maxRPM int) *client[T] {
	v := &client[T]{adapter: new(GenericAdapter[T])}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	v.client = &http.Client{Transport: tr}
	v.ratelimiter = limiter.NewLimiter(maxRPM)

	return v
}

func (c *client[T]) do(req *http.Request) (*http.Response, error) {
	var resp *http.Response

	err := c.ratelimiter.Call(func() error {
		var er error
		resp, er = c.client.Do(req)
		if er != nil {
			return er
		}

		return nil
	})

	return resp, err
}

func (c *client[T]) getBody(resp *http.Response) ([]byte, error) {
	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrNotFound

	case http.StatusTooManyRequests:
		// c.stde.Printf("too many requests")
		return nil, ErrTooManyRequests

	case http.StatusOK:
		scanner := bufio.NewScanner(resp.Body)

		var response string

		for scanner.Scan() {
			response += scanner.Text()
		}

		return []byte(response), nil

	default:
		scanner := bufio.NewScanner(resp.Body)

		var er string

		for scanner.Scan() {
			er += scanner.Text()
		}
		// c.stde.Printf("error code %d: url:%s %s", resp.StatusCode, resp.Request.URL.RequestURI(), er)

		return []byte(er), ErrUnknownError
	}
}

func (c *client[T]) executeWithContext(
	ctx context.Context,
	method,
	url string,
	body io.Reader,
	headers map[string]string,
) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	var res *http.Response

	if c.ratelimiter != nil {
		res, err = c.do(req)
	} else {
		res, err = c.client.Do(req)
	}

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseBody, err := c.getBody(res)
	if err != nil {
		return nil, fmt.Errorf("error status: %d; %s", res.StatusCode, responseBody)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status: %d %s; error: %s", res.StatusCode, body, err.Error())
	}

	parsed, err := c.adapter.ParseResponse(responseBody)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}

	return parsed, nil
}

func (c *client[T]) Execute(method, url string, body io.Reader, headers map[string]string) (*T, error) {
	ctx := context.Background()

	return c.executeWithContext(ctx, method, url, body, headers)
}

func (c *client[T]) ExecuteWithContext(
	ctx context.Context,
	method,
	url string,
	body io.Reader,
	headers map[string]string,
) (*T, error) {
	return c.executeWithContext(ctx, method, url, body, headers)
}
