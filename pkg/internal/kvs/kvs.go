package kvs

type KVS interface {
	Get(k string) ([]byte, error)
	Put(k string, v []byte) error
	Del(k string) error
}
