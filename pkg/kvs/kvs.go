package kvs

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrEmptyKey    = errors.New("key is empty")
)

type KVS interface {
	Get(k string) ([]byte, error)
	Put(k string, v []byte) error
	Del(k string) error
}

type ServerImpl struct {
	server    *grpc.Server
	lis       *net.Listener
	kvsClient KVS
}

type routeKVSClientService struct {
	kvsClient KVS
}

func (r *routeKVSClientService) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	if len(in.Key) == 0 {
		return nil, ErrEmptyKey
	}

	b, err := r.kvsClient.Get(in.Key)
	if err != nil {
		return nil, err
	}

	res := GetResponse{Key: in.Key, Values: b}

	return &res, nil
}

func (r *routeKVSClientService) Del(ctx context.Context, in *DelRequest) (*Response, error) {
	if len(in.Key) == 0 {
		return nil, ErrEmptyKey
	}

	err := r.kvsClient.Del(in.Key)
	if err != nil {
		return nil, err
	}

	res := Response{Code: 0}

	return &res, nil
}

func (r *routeKVSClientService) Put(ctx context.Context, in *PutRequest) (*Response, error) {
	if len(in.Key) == 0 {
		return nil, ErrEmptyKey
	}

	err := r.kvsClient.Put(in.Key, in.Value)
	if err != nil {
		return nil, err
	}

	res := Response{Code: 0}

	return &res, nil
}

func NewKVS(kvsClient KVS, port int64) (*ServerImpl, error) {
	s := ServerImpl{}
	s.kvsClient = kvsClient

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)

		return nil, err
	}

	s.lis = &lis

	var opts []grpc.ServerOption
	s.server = grpc.NewServer(opts...)
	router := routeKVSClientService{kvsClient: kvsClient}
	RegisterKVSServiceServer(s.server, &router)

	return &s, nil
}

func (s *ServerImpl) Listen() {
	if err := s.server.Serve(*s.lis); err != nil {
		return
	}
}
