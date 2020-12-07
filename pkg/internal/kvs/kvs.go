package kvs

import (
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

func NewKVS(port int64) (*ServerImpl, error) {
	s := ServerImpl{}

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
