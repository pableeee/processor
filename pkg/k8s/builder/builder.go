package builder

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pableeee/processor/pkg/internal/kvs"
	"github.com/pableeee/processor/pkg/internal/lock"
	"github.com/pableeee/processor/pkg/k8s/provider"
	"github.com/pableeee/processor/pkg/k8s/provider/types"
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

type Model struct {
	// Project name
	Project string

	// Github repo url
	Repo string

	//
	URL string

	//
	ServivceName string

	//
	Type types.ServiceType
}

type Services interface {
	Get(k string) ([]byte, error)
}

type Builder struct {
	// Configuration data

	// Service map
	services kvs.KVS

	// Distributed Locks
	lck lock.Client

	// Client to build infra
	provider provider.InfraProvider
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) GetServices() Services {
	return b.services
}

func (b *Builder) WithProvider(p provider.InfraProvider) *Builder {
	b.provider = p

	return b
}

func (b *Builder) WithLock(l lock.Client) *Builder {
	b.lck = l

	return b
}

func (b *Builder) WithKVS(k kvs.KVS) *Builder {
	b.services = k

	return b
}

func (b *Builder) BuildKVS(m Model) error {
	return b.buildService(m.Project, m.ServivceName, m.Type,
		func() error {
			return b.provider.BuildKVS(m.ServivceName)
		})
}

func (b *Builder) BuildLock(m Model) error {
	return b.buildService(m.Project, m.ServivceName, m.Type,
		func() error {
			return b.provider.BuildLock(m.ServivceName)
		})
}

func (b *Builder) BuildQueue(m Model) error {
	return b.buildService(m.Project, m.ServivceName, m.Type,
		func() error {
			return b.provider.BuildQueue(m.ServivceName)
		})
}

func (b *Builder) buildService(ns, id string, t types.ServiceType,
	f func() error) error {
	if len(id) < 3 {
		// Ids mix 3 char
		return fmt.Errorf("invalid argument")
	}

	key := fmt.Sprintf("kvs:%s", id)

	l, err := b.lck.Lock(key, 50)
	if err != nil {
		return fmt.Errorf("failed locking id: %w", err)
	}

	defer func() {
		err = b.lck.Unlock(l)
		if err != nil {
			log.Println("incrementar metrica")
		}
	}()

	err = f()
	if err != nil {
		return fmt.Errorf("failed building kvs : %w", err)
	}

	s := Service{
		Name: fmt.Sprintf("%s_%s_%s", t, ns, id),
		Type: t,
	}

	bts, err := json.Marshal(s)
	if err != nil {
		// TODO borrar la infra creada?
		return fmt.Errorf("failed marshaling: %w", err)
	}

	err = b.services.Put(s.Name, bts)
	if err != nil {
		// TODO borrar la infra creada?
		return fmt.Errorf("failed updating kvs: %w", err)
	}

	return nil
}
