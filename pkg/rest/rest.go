package rest

import (
	"bufio"
	"context"
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

type Client struct {
	ratelimiter limiter.RateLimiter
	client      *http.Client
}

func NewRestClient() *Client {
	v := &Client{}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	v.client = &http.Client{Transport: tr}

	return v
}
func WithLimiter(maxRPM int) *Client {
	v := &Client{}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	v.client = &http.Client{Transport: tr}
	v.ratelimiter = limiter.NewLimiter(maxRPM)

	return v
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
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

func (c *Client) getBody(resp *http.Response) (string, error) {
	switch resp.StatusCode {
	case http.StatusNotFound:
		return "", ErrNotFound

	case http.StatusTooManyRequests:
		// c.stde.Printf("too many requests")
		return "", ErrTooManyRequests

	case http.StatusOK:
		scanner := bufio.NewScanner(resp.Body)

		var response string

		for scanner.Scan() {
			response += scanner.Text()
		}

		return response, nil

	default:
		scanner := bufio.NewScanner(resp.Body)

		var er string

		for scanner.Scan() {
			er += scanner.Text()
		}
		// c.stde.Printf("error code %d: url:%s %s", resp.StatusCode, resp.Request.URL.RequestURI(), er)

		return er, ErrUnknownError
	}
}

func (c *Client) Execute(method, url string, body io.Reader, headers map[string]string) (string, error) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return "", err
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
		return "", err
	}
	defer res.Body.Close()

	responseBody, err := c.getBody(res)
	if err != nil {
		return "", fmt.Errorf("error status: %d; %s", res.StatusCode, responseBody)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error status: %d %s; error: %s", res.StatusCode, body, err.Error())
	}

	return responseBody, nil
}
