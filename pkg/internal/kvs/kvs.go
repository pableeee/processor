package kvs

import "errors"

var (
	KeyNotFound = errors.New("key not found")
)

type KVS interface {
	Get(k string) ([]byte, error)
	Put(k string, v []byte) error
	Del(k string) error
}
