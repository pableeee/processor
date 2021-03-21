package kvs

import (
	"errors"

	"github.com/go-redis/redis"
	rds "github.com/pableeee/processor/pkg/internal/redis"
)

type RedisKVS struct {
	client *redis.Client
}

func MakeRedisKVS() (*RedisKVS, error) {
	r := &RedisKVS{}

	client, err := rds.MakeRedisClient("localhost", 6379)
	if err != nil {
		return nil, err
	}

	r.client = client

	return r, nil
}

func (r *RedisKVS) Get(k string) ([]byte, error) {
	val, err := r.client.Get(k).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, ErrKeyNotFound
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
		return ErrKeyNotFound
	} else if err != nil {
		return err
	}

	return nil
}
