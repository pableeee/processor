package kvs

import (
	"fmt"
	"sync"
)

var instance *LocalKVS

// LocalInfra for local tests
type LocalKVS struct {
	db  *map[string][]byte
	mux *sync.Mutex
}

// Get implements a local get service
func (infra *LocalKVS) Get(k string) ([]byte, error) {
	if len(k) == 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	infra.mux.Lock()
	s, found := (*infra.db)[k]
	infra.mux.Unlock()

	if !found {
		return nil, fmt.Errorf("key %s not found", k)
	}

	return s, nil

}

// Put implements a o local put service
func (infra *LocalKVS) Put(k string, v []byte) error {
	if len(k) == 0 {
		return fmt.Errorf("invalid user id")
	}

	infra.mux.Lock()
	(*infra.db)[k] = v
	infra.mux.Unlock()

	return nil
}

// MakeLocalInfra makes an instances of a local infra
func MakeLocalKVS() *LocalKVS {
	if instance == nil {
		instance = new(LocalKVS)
		instance.mux = &sync.Mutex{}
		m := make(map[string][]byte)
		instance.db = &m
	}

	return instance
}
