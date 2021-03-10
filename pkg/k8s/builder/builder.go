package builder

import (
	"fmt"

	"github.com/alexellis/arkade/pkg/apps"
	"github.com/alexellis/arkade/pkg/types"
)

type InfraProvider interface {
	BuildKVS(id string) error
	BuildLock(id string) error
	BuildQueue(id string) error
}

func NewInfraProvider() InfraProvider {
	return &defaultInfraProvider{}
}

type defaultInfraProvider struct {
}

func (d *defaultInfraProvider) BuildQueue(id string) error {
	return BuildRedisNats()
}

func (d *defaultInfraProvider) BuildKVS(id string) error {
	return BuildRedisKVS()
}
func (d *defaultInfraProvider) BuildLock(id string) error {
	return BuildRedisKVS()
}

func BuildRedisKVS() error {
	redisAppOptions := types.DefaultInstallOptions().
		WithNamespace("namespace").
		WithHelmRepo("bitnami-redis/redis").
		WithHelmURL("https://charts.bitnami.com/bitnami").
		WithOverrides(nil).
		WithWait(false).
		WithHelmUpdateRepo(false).
		WithKubeconfigPath("kubeConfigPath")

	_, err := apps.MakeInstallChart(redisAppOptions)
	if err != nil {
		return fmt.Errorf("fail applying redis chart: %v", err)
	}

	return nil
}

func BuildRedisNats() error {
	redisAppOptions := types.DefaultInstallOptions().
		WithNamespace("namespace").
		WithHelmRepo("nats/nats").
		WithHelmURL("https://nats-io.github.io/k8s/helm/charts/").
		WithOverrides(nil).
		WithWait(false).
		WithHelmUpdateRepo(false).
		WithKubeconfigPath("kubeConfigPath")

	_, err := apps.MakeInstallChart(redisAppOptions)
	if err != nil {
		return fmt.Errorf("fail applying nats chart: %v", err)
	}

	return nil
}
