package kvs

import (
	"github.com/pableeee/processor/pkg/internal/kvs"
	"google.golang.org/grpc"
)

type Service interface {
	Listen()
}

const defaultPort = 33333

func NewLocalKVS() (Service, error) {
	kvsClient := &kvs.LocalKVS{}
	s, err := kvs.NewKVS(kvsClient, defaultPort)

	return s, err
}

func NewRedisKVS(port int64) (Service, error) {
	kvsClient := kvs.MakeRedisKVS()
	s, err := kvs.NewKVS(kvsClient, defaultPort)

	return s, err
}

type kvsGRPCClient struct {
}

func NewKVSClient(address string, port int64) error {
	var opts []grpc.DialOption

	_, err := grpc.Dial(address, opts...)
	if err != nil {
		return err
	}

	return nil
}
