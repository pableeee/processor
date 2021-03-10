package types

type ServiceType string

const (
	KVS   ServiceType = "kvs"
	Lock  ServiceType = "lock"
	Queue ServiceType = "queue"
)
