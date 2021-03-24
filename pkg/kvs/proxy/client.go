package proxy

import (
	"context"
	"fmt"

	"github.com/pableeee/processor/pkg/kvs"
	"google.golang.org/grpc"
)

type Client struct {
	client kvs.KVSServiceClient
}

func NewClient(address string) *Client {
	con, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil
	}

	return &Client{client: kvs.NewKVSServiceClient(con)}
}

func (p *Client) Get(k string) ([]byte, error) {
	ctx := context.Background()

	r, err := p.client.Get(ctx, &kvs.GetRequest{Key: k})
	if err != nil {
		return nil, fmt.Errorf("failed getting key %s %w", k, err)
	}

	if r.Error != nil {
		return nil, fmt.Errorf("failed getting key %s %s", k, r.Error.Error)
	}

	return r.Values, nil
}

func (p *Client) Put(k string, v []byte) error {
	ctx := context.Background()

	r, err := p.client.Put(ctx, &kvs.PutRequest{Key: k, Value: v})
	if err != nil {
		return fmt.Errorf("failed getting key %s %w", k, err)
	}

	if r.Error != "" {
		return fmt.Errorf("failed getting key %s %s", k, r.Code)
	}

	return nil
}

func (p *Client) Del(k string) error {
	ctx := context.Background()

	r, err := p.client.Del(ctx, &kvs.DelRequest{Key: k})
	if err != nil {
		return fmt.Errorf("failed getting key %s %w", k, err)
	}

	if r.Error != "" {
		return fmt.Errorf("failed getting key %s %s", k, r.Code)
	}

	return nil
}
