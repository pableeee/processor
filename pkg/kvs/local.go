package kvs

import (
	"fmt"
	"log"
	"sync"
)

var ErrorInvalidID = fmt.Errorf("invalid user id")

// LocalKVS for local tests.
type LocalKVS struct {
	db  *map[string][]byte
	mux *sync.Mutex
}

func NewLocal() KVS {
	k := LocalKVS{}
	m := make(map[string][]byte)
	k.db = &m
	k.mux = &sync.Mutex{}

	return &k
}

// Get implements a local get service.
func (infra *LocalKVS) Get(k string) ([]byte, error) {
	if len(k) == 0 {
		return nil, ErrorInvalidID
	}

	infra.mux.Lock()
	s, found := (*infra.db)[k]
	infra.mux.Unlock()

	if !found {
		return nil, ErrKeyNotFound
	}

	return s, nil
}

// Put implements a o local put service.
func (infra *LocalKVS) Put(k string, v []byte) error {
	if len(k) == 0 {
		return ErrorInvalidID
	}

	infra.mux.Lock()
	(*infra.db)[k] = v
	infra.mux.Unlock()

	return nil
}

// Put implements a o local put service.
func (infra *LocalKVS) Del(k string) error {
	if len(k) == 0 {
		return ErrorInvalidID
	}

	infra.mux.Lock()
	defer infra.mux.Unlock()

	_, found := (*infra.db)[k]
	if !found {
		log.Printf("could not delete game:%s", k)

		return nil
	}

	delete(*infra.db, k)

	return nil
}

// MakeLocalInfra makes an instances of a local infra.
func MakeLocalKVS() *LocalKVS {
	instance := new(LocalKVS)
	instance.mux = &sync.Mutex{}
	m := make(map[string][]byte)
	instance.db = &m

	return instance
}
