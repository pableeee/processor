package proxy

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/pableeee/processor/pkg/kvs"
	"google.golang.org/grpc"
)

type RPCServer interface {
	Serve(lis net.Listener) error
	Stop()
}

type Server struct {
	client kvs.KVS
	serv   RPCServer
	lis    net.Listener
}

func (s *Server) Get(c context.Context, r *kvs.GetRequest) (*kvs.GetResponse, error) {
	b, err := s.client.Get(r.Key)
	if errors.Is(err, kvs.ErrKeyNotFound) {
		return &kvs.GetResponse{Error: &kvs.Response{Error: err.Error()}}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed getting int kvs %s: %w", r.Key, err)
	}

	return &kvs.GetResponse{Key: r.Key, Values: b}, nil
}

func (s *Server) Put(c context.Context, r *kvs.PutRequest) (*kvs.Response, error) {
	if err := s.client.Put(r.Key, r.Value); err != nil {
		return nil, fmt.Errorf("failed saving in kvs %s: %w", r.Key, err)
	}

	return &kvs.Response{}, nil
}

func (s *Server) Del(c context.Context, r *kvs.DelRequest) (*kvs.Response, error) {
	if err := s.client.Del(r.Key); errors.Is(err, kvs.ErrKeyNotFound) {
		return &kvs.Response{Error: err.Error()}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed deletig in kvs %s: %w", r.Key, err)
	}

	return &kvs.Response{}, nil
}

func NewServer(addr string) *Server {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sv := &Server{
		serv: s,
		lis:  lis,
	}

	kvs.RegisterKVSServiceServer(s, sv)

	return sv
}

func (s *Server) WithClient(c kvs.KVS) *Server {
	s.client = c

	return s
}

func (s *Server) Serve() error {
	if err := s.serv.Serve(s.lis); err != nil {
		return fmt.Errorf("failed starting grpc server %w", err)
	}

	return nil
}

func (s *Server) Close() {
	s.serv.Stop()
}
