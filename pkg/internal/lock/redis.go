package lock

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	rds "github.com/pableeee/processor/pkg/internal/redis"
)

type redisLockClient struct {
	client *redis.Client
}

func MakeRedisLock() (Client, error) {
	r := redisLockClient{}

	client, err := rds.MakeRedisClient("localhost", 6379)
	if err != nil {
		return nil, err
	}

	r.client = client

	return &r, nil
}

func (l *redisLockClient) Get(resource string) (Lock, error) {
	lck, err := l.retrieveLock(resource)
	if err != nil {
		return nil, err
	}

	return lck, nil
}

func (l *redisLockClient) Lock(resource string, ttl int) (Lock, error) {
	lck := NewLock(resource, ttl)

	ok, err := l.client.SetNX(resource, lck, time.Duration(ttl)*(time.Millisecond)).Result()
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrLocked
	}

	return lck, nil
}

func (l *redisLockClient) unmarshalResponse(payload string) (Lock, error) {
	var lck lock

	err := json.Unmarshal([]byte(payload), &lck)
	if err != nil {
		return nil, ErrInternal
	}

	return &lck, nil
}

func (l *redisLockClient) retrieveLock(key string) (Lock, error) {
	r, err := l.client.Get(key).Result()
	if err != nil {
		return nil, ErrNotFound
	}

	lck, err := l.unmarshalResponse(r)
	if err != nil {
		return nil, err
	}

	return lck, nil
}

func (l *redisLockClient) KeepAlive(lock Lock) (Lock, error) {
	lck, err := l.retrieveLock(lock.GetResource())
	if err != nil {
		return nil, err
	}

	if lock.GetToken() != lck.GetToken() {
		return nil, ErrLocked
	}

	ok, err := l.client.Expire(lock.GetResource(), time.Duration(lock.GetTTL())*time.Millisecond).Result()
	if err != nil || !ok {
		return nil, ErrInternal
	}

	return lck, nil
}

func (l *redisLockClient) Unlock(lock Lock) error {
	lck, err := l.retrieveLock(lock.GetResource())
	if err != nil {
		return err
	}

	if lock.GetToken() != lck.GetToken() {
		return ErrLocked
	}

	_, err = l.client.Del(lck.GetResource()).Result()
	if err != nil {
		return ErrInternal
	}

	return nil
}
