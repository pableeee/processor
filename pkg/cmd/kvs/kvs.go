package kvs

import (
	"context"
	"fmt"

	ukvs "github.com/pableeee/processor/pkg/internal/kvs"
	"google.golang.org/grpc"
)

type Service interface {
	Listen()
}

const defaultPort = 33333

func NewLocalKVS() (Service, error) {
	kvsClient := ukvs.MakeLocalKVS()
	s, err := ukvs.NewKVS(kvsClient, defaultPort)

	return s, err
}

func NewRedisKVS(port int64) (Service, error) {
	kvsClient, err := ukvs.MakeRedisKVS()
	if err != nil {
		return nil, err
	}

	s, err := ukvs.NewKVS(kvsClient, defaultPort)

	return s, err
}

type grpcKVSClient struct {
	client ukvs.KVSServiceClient
}

func (c *grpcKVSClient) Get(k string) ([]byte, error) {
	req := ukvs.GetRequest{Key: k}

	res, err := c.client.Get(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return res.Values, nil
}

func (c *grpcKVSClient) Put(k string, v []byte) error {
	req := ukvs.PutRequest{Key: k, Value: v}

	_, err := c.client.Put(context.Background(), &req)
	if err != nil {
		return err
	}

	// TODO: check for application level error
	return nil
}

func (c *grpcKVSClient) Del(k string) error {
	req := ukvs.DelRequest{Key: k}

	_, err := c.client.Del(context.Background(), &req)
	if err != nil {
		return err
	}

	// TODO: check for application level error
	return nil
}

// NewKVSClient create a new grpc kvs client, and returns it as an kvs interface.
func NewKVSClient(address string, port int64) (ukvs.KVS, error) {
	c, err := grpc.Dial(fmt.Sprintf("%s:%d", address, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcKVSClient{client: ukvs.NewKVSServiceClient(c)}, nil
}
