package proxy

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/pableeee/processor/pkg/kvs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockServer struct {
	mock.Mock
}

func (m *mockServer) Serve(lis net.Listener) error {
	args := m.Called(lis)

	return args.Error(0)
}
func (m *mockServer) Stop() {
	return
}

type mockListener struct {
	mock.Mock
}

func (m *mockListener) Accept() (net.Conn, error) {
	args := m.Called()
	if b, ok := args.Get(0).(net.Conn); ok {
		return b, args.Error(1)
	}

	return nil, args.Error(1)
}
func (m *mockListener) Close() error {
	args := m.Called()

	return args.Error(1)
}
func (m *mockListener) Addr() net.Addr {
	args := m.Called()
	if b, ok := args.Get(0).(net.Addr); ok {
		return b
	}

	return nil
}

type mockKVS struct {
	mock.Mock
}

func (m *mockKVS) Get(k string) ([]byte, error) {
	args := m.Called(k)

	if b, ok := args.Get(0).([]byte); ok {
		return b, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockKVS) Put(k string, v []byte) error {
	args := m.Called(k, v)

	return args.Error(0)
}
func (m *mockKVS) Del(k string) error {
	args := m.Called(k)

	return args.Error(0)
}

func TestGet_KeyNotFound(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Get", "123456").Return(nil, kvs.ErrKeyNotFound)
	r, err := s.Get(context.Background(), &kvs.GetRequest{Key: "123456"})

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Contains(t, r.Error.Error, kvs.ErrKeyNotFound.Error())

}

func TestGet_Error(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Get", "123456").Return(nil, errors.New("some error"))
	r, err := s.Get(context.Background(), &kvs.GetRequest{Key: "123456"})

	assert.NotNil(t, err)
	assert.Nil(t, r)
	assert.Contains(t, err.Error(), "some error")

}

func TestGet_OK(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Get", "123456").Return(
		[]byte(`some response`), nil)
	r, err := s.Get(context.Background(), &kvs.GetRequest{Key: "123456"})

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, r.Values, []byte(`some response`))

}

func TestPut_Fail(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Put", "123456", []byte(`value`)).Return(errors.New("some error"))
	r, err := s.Put(context.Background(),
		&kvs.PutRequest{Key: "123456",
			Value: []byte(`value`),
		})

	assert.NotNil(t, err)
	assert.Nil(t, r)
	assert.Contains(t, err.Error(), "some error")

}

func TestPut_OK(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Put", "123456", []byte(`value`)).Return(nil)
	r, err := s.Put(context.Background(),
		&kvs.PutRequest{
			Key:   "123456",
			Value: []byte(`value`),
		})

	assert.Nil(t, err)
	assert.NotNil(t, r)

}
func TestDel_KeyNotFound(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Del", "123456").Return(kvs.ErrKeyNotFound)
	r, err := s.Del(context.Background(), &kvs.DelRequest{Key: "123456"})

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Contains(t, r.Error, kvs.ErrKeyNotFound.Error())
}

func TestDel_Error(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Del", "123456").Return(errors.New("some error"))
	r, err := s.Del(context.Background(), &kvs.DelRequest{Key: "123456"})

	assert.NotNil(t, err)
	assert.Nil(t, r)
	assert.Contains(t, err.Error(), "some error")

}

func TestDel_OK(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	k.On("Del", "123456").Return(nil)
	r, err := s.Del(context.Background(),
		&kvs.DelRequest{
			Key: "123456",
		})

	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestServe_OK(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	sv.On("Serve", mock.Anything).Return(nil)

	err := s.Serve()
	assert.Nil(t, err)
}

func TestServe_Fail(t *testing.T) {
	k := new(mockKVS)
	lis := new(mockListener)
	sv := new(mockServer)
	s := &Server{
		client: k,
		serv:   sv,
		lis:    lis,
	}

	sv.On("Serve", mock.Anything).Return(errors.New("some error"))
	err := s.Serve()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed starting grpc server")
}
