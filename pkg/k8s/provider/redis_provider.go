package provider

import (
	"fmt"

	"github.com/alexellis/arkade/pkg/apps"
	"github.com/alexellis/arkade/pkg/types"
)

type Model struct {
	// Project name
	Name string

	// Github repo url
	Repo string

	//
	URL string

	Port int
}

func NewInfraProvider() InfraProvider {
	return &defaultInfraProvider{}
}

type defaultInfraProvider struct {
}

func (d *defaultInfraProvider) install(cfg string, data Model) error {
	options := types.DefaultInstallOptions().
		WithNamespace(data.Name).
		WithHelmRepo(data.Repo).
		WithHelmURL(data.URL).
		WithOverrides(nil).
		WithWait(false).
		WithHelmUpdateRepo(false).
		WithKubeconfigPath(cfg)

	_, err := apps.MakeInstallChart(options)
	if err != nil {
		return fmt.Errorf("fail applying redis chart: %v", err)
	}

	return nil
}

func (d *defaultInfraProvider) BuildLock(id string) error {
	err := d.install("",
		Model{
			Name: id,
			Repo: "bitnami-redis/redis",
			URL:  "https://charts.bitnami.com/bitnami",
			// cambio el puerto
			Port: 6389,
		})

	if err != nil {
		return fmt.Errorf("fail applying redis chart: %v", err)
	}

	return nil
}

func (d *defaultInfraProvider) BuildKVS(id string) error {
	err := d.install("",
		Model{
			Name: id,
			Repo: "bitnami-redis/redis",
			URL:  "https://charts.bitnami.com/bitnami",
			Port: 6379,
		})

	if err != nil {
		return fmt.Errorf("fail applying redis chart: %v", err)
	}

	return nil
}

func (d *defaultInfraProvider) BuildQueue(id string) error {
	err := d.install("", Model{
		Name: id,
		Repo: "https://nats-io.github.io/k8s/helm/charts/",
		URL:  "nats/nats",
		Port: 4444,
	})
	if err != nil {
		return fmt.Errorf("fail applying redis chart: %v", err)
	}

	return nil
}
