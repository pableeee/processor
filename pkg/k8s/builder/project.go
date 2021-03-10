package builder

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pableeee/processor/pkg/internal/kvs"
	"github.com/pableeee/processor/pkg/internal/lock"
	"github.com/pableeee/processor/pkg/k8s/builder/types"
)

type Implementation struct {
	Repo string
	URL  string
}

type Configuration struct {
	URL  string
	Port int
}

type Service struct {
	Name   string
	Type   types.ServiceType
	Chart  Implementation
	Config Configuration
}

type ProjectData struct {
	// Project name
	Name string

	// Github repo url
	repo string
}

type Project struct {
	// Configuration data
	Data ProjectData

	// Service map
	services kvs.KVS

	// Distributed Locks
	lck lock.Client

	// Client to build infra
	builder InfraProvider
}

func (p *Project) BuildKVS(id string) error {
	return p.buildService(id,
		types.KVS,
		Configuration{URL: "localhost", Port: 6379},
		Implementation{},
		func(i string) error {
			return p.builder.BuildKVS(id)
		})
}

func (p *Project) BuildLock(id string) error {
	return p.buildService(id,
		types.KVS,
		Configuration{URL: "localhost", Port: 6379},
		Implementation{},
		func(i string) error {
			return p.builder.BuildLock(id)
		})
}

func (p *Project) BuildQueue(id string) error {
	return p.buildService(id,
		types.KVS,
		Configuration{URL: "localhost", Port: 5555},
		Implementation{},
		func(i string) error {
			return p.builder.BuildQueue(id)
		})
}
func (p *Project) buildService(id string, t types.ServiceType, cfg Configuration, chart Implementation,
	f func(i string) error) error {
	if len(id) < 3 {
		//Ids mix 3 char
		return fmt.Errorf("invalid argument")
	}

	key := fmt.Sprintf("kvs:%s", id)

	l, err := p.lck.Lock(key, 50)
	if err != nil {
		return fmt.Errorf("failed locking id: %w", err)
	}

	defer func() {
		err = p.lck.Unlock(l)
		if err != nil {
			log.Println("incrementar metrica")
		}
	}()

	err = f(id)
	if err != nil {
		return fmt.Errorf("failed building kvs : %w", err)
	}

	s := Service{
		Name:   fmt.Sprintf("%s_%s_%s", t, p.Data.Name, id),
		Type:   t,
		Config: cfg,
		Chart:  chart,
	}

	b, err := json.Marshal(s)
	if err != nil {
		// TODO borrar la infra creada?
		return fmt.Errorf("failed marshaling: %w", err)
	}

	err = p.services.Put(s.Name, b)
	if err != nil {
		// TODO borrar la infra creada?
		return fmt.Errorf("failed updating kvs: %w", err)
	}

	return nil
}
