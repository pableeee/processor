package infra

import (
	"fmt"
	"sync"
)

var instance *LocalInfra

// LocalInfra for local tests
type LocalInfra struct {
	db  *map[string][]Server
	mux *sync.Mutex
}

// Get implements a local get service
func (infra *LocalInfra) Get(userID string) ([]Server, error) {
	if len(userID) == 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	infra.mux.Lock()
	s, found := (*infra.db)[userID]
	infra.mux.Unlock()

	if !found {
		return []Server{}, nil
	}

	return s, nil

}

// Put implements a o local put service
func (infra *LocalInfra) Put(userID string, s Server) error {
	if len(userID) == 0 {
		return fmt.Errorf("invalid user id")
	}

	infra.mux.Lock()
	usr, found := (*infra.db)[userID]
	if !found {
		usr = make([]Server, 0)
	}

	infra.mux.Unlock()

	usr = append(usr, s)
	(*infra.db)[userID] = usr

	return nil
}

// MakeLocalInfra makes an instances of a local infra
func MakeLocalInfra() *LocalInfra {
	if instance == nil {
		instance = new(LocalInfra)
		instance.mux = &sync.Mutex{}
		m := make(map[string][]Server)
		instance.db = &m
	}

	return instance
}
