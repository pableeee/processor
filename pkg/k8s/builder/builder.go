package builder

import (
	"fmt"

	"github.com/pableeee/processor/pkg/k8s/provider"
	"github.com/pableeee/processor/pkg/k8s/provider/types"
	"github.com/pableeee/processor/pkg/repository"
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
	repo repository.Repository

	// Client to build infra
	provider provider.InfraProvider
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) GetService(proj, name string, t types.ServiceType) (Model, error) {
	var m Model

	key := fmt.Sprintf("%s:%s_%s", proj, t, name)
	if err := b.repo.Get(key, &m); err != nil {
		return m, fmt.Errorf("failed retrieving service: %s %w", key, err)
	}

	return m, nil
}

func (b *Builder) GetProyect(id string) (Model, error) {
	var m Model

	err := b.repo.Get(id, &m)
	if err != nil {
		return m, fmt.Errorf("failed getting project: %s %w", id, err)
	}

	return m, nil
}

func (b *Builder) WithProvider(p provider.InfraProvider) *Builder {
	b.provider = p

	return b
}

func (b *Builder) WithRepository(r repository.Repository) *Builder {
	b.repo = r

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

	if err := f(); err != nil {
		return fmt.Errorf("failed building kvs : %w", err)
	}

	s := Service{
		Name: fmt.Sprintf("%s:%s_%s", ns, t, id),
		Type: t,
	}

	if err := b.repo.Save(s.Name, s); err != nil {
		// TODO borrar la infra creada?
		return fmt.Errorf("failed saving service: %w", err)
	}

	return nil
}
