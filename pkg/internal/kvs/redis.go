package kvs

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var ErrorErrKeyNotFound = fmt.Errorf("key not found")

type RedisKVS struct {
	client *redis.Client
}

func MakeRedisKVS() *RedisKVS {
	r := &RedisKVS{}

	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return r
}

func (r *RedisKVS) Get(k string) ([]byte, error) {
	val, err := r.client.Get(k).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, ErrorErrKeyNotFound
	} else if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisKVS) Put(k string, v []byte) error {
	res := r.client.Set(k, v, 0).Err()

	return res
}

func (r *RedisKVS) Del(k string) error {
	_, err := r.client.Del(k).Result()
	if errors.Is(err, redis.Nil) {
		return ErrorErrKeyNotFound
	} else if err != nil {
		return err
	}

	return nil
}
