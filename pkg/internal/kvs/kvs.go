package kvs

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var ErrKeyNotFound = errors.New("key not found")

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

func (r *routeKVSClientService) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	if len(in.Key) <= 0 {
		return nil, fmt.Errorf("key is empty")
	}
	
	b, err := r.kvsClient.Get(in.Key)
	if err != nil {
		return nil, err
	}

	GetResponse{Key: in.Key, Values: }
}

func (r *routeKVSClientService) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Response, error) {

}

func (r *routeKVSClientService) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error) {

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

	return &s, nil
}

func (s *ServerImpl) Listen() {
	s.server.Serve(*s.lis)
}
