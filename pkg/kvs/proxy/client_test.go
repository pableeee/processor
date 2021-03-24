package proxy

import (
	"context"
	"errors"
	"testing"

	"github.com/pableeee/processor/pkg/kvs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type mockKVSServiceClient struct {
	mock.Mock
}

func (m *mockKVSServiceClient) Get(ctx context.Context, in *kvs.GetRequest, opts ...grpc.CallOption) (*kvs.GetResponse, error) {
	args := m.Called(ctx, in, opts)
	if b, ok := args.Get(0).(*kvs.GetResponse); ok {
		return b, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *mockKVSServiceClient) Put(ctx context.Context, in *kvs.PutRequest, opts ...grpc.CallOption) (*kvs.Response, error) {
	args := m.Called(ctx, in, opts)
	if b, ok := args.Get(0).(*kvs.Response); ok {
		return b, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *mockKVSServiceClient) Del(ctx context.Context, in *kvs.DelRequest, opts ...grpc.CallOption) (*kvs.Response, error) {
	args := m.Called(ctx, in, opts)
	if b, ok := args.Get(0).(*kvs.Response); ok {
		return b, args.Error(1)
	}

	return nil, args.Error(1)
}

func TestGet_Ok(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Get", context.Background(), &kvs.GetRequest{Key: "123456"}, par).
		Return(&kvs.GetResponse{
			Key:    "123456",
			Values: []byte(`value`)}, nil)

	b, err := s.Get("123456")
	assert.Nil(t, err)
	assert.NotNil(t, b)
	assert.Equal(t, string(b), "value")
}

func TestGet_FailLocal(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Get", context.Background(), &kvs.GetRequest{Key: "123456"}, par).
		Return(nil, errors.New("some local error"))

	b, err := s.Get("123456")
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed getting key")
}

func TestGet_FailRemote(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Get", context.Background(), &kvs.GetRequest{Key: "123456"}, par).
		Return(&kvs.GetResponse{Error: &kvs.Response{Error: "some remote error"}}, nil)

	b, err := s.Get("123456")
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "some remote error")
}

func TestPut_Ok(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Put", context.Background(), &kvs.PutRequest{Key: "123456", Value: []byte(`value`)}, par).
		Return(&kvs.Response{
			Error: ,
		}, nil)

	b, err := s.Get("123456")
	assert.Nil(t, err)
	assert.NotNil(t, b)
	assert.Equal(t, string(b), "value")
}

func TestGet_FailLocal(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Get", context.Background(), &kvs.GetRequest{Key: "123456"}, par).
		Return(nil, errors.New("some local error"))

	b, err := s.Get("123456")
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed getting key")
}

func TestGet_FailRemote(t *testing.T) {
	k := new(mockKVSServiceClient)
	s := &Client{
		client: k,
	}
	var par []grpc.CallOption

	k.On("Get", context.Background(), &kvs.GetRequest{Key: "123456"}, par).
		Return(&kvs.GetResponse{Error: &kvs.Response{Error: "some remote error"}}, nil)

	b, err := s.Get("123456")
	assert.Nil(t, b)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "some remote error")
}
