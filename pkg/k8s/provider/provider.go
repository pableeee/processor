package provider

type InfraProvider interface {
	BuildKVS(id string) error
	BuildLock(id string) error
	BuildQueue(id string) error
}
