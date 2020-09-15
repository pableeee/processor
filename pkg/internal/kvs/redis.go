package kvs

import (
	"fmt"

	"github.com/go-redis/redis"
)

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
	if err == redis.Nil {
		return nil, fmt.Errorf("%s does not exist", k)
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
	if err == redis.Nil {
		return fmt.Errorf("%s does not exist", k)
	} else if err != nil {
		return err
	}

	return nil
}
