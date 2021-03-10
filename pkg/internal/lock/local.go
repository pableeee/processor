package lock

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type localLockClient struct {
	mux   *sync.RWMutex
	locks map[string]lock
}

func NewLocal() Client {
	lck := localLockClient{}
	lck.mux = &sync.RWMutex{}
	lck.locks = make(map[string]lock)

	return &lck
}

func (l *localLockClient) Get(resource string) (Lock, error) {
	if len(resource) == 0 {
		return nil, ErrInvalidArg
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	lck, found := l.locks[resource]
	if !found {
		return nil, ErrNotFound
	}

	if lck.expired() {
		delete(l.locks, resource)

		return nil, ErrExpired
	}

	return &lck, nil
}

func (l *localLockClient) Lock(resource string, ttl int) (Lock, error) {
	if len(resource) == 0 || ttl < 0 {
		return nil, ErrInvalidArg
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	lck, found := l.locks[resource]
	if found && !lck.expired() {
		return nil, ErrLocked
	}

	token := uuid.New()
	t := time.Now().Add(time.Duration(ttl) * time.Second)
	lck = lock{
		TTL:        ttl,
		Resource:   resource,
		Origin:     "",
		Token:      token.String(),
		expiration: &t,
	}

	l.locks[resource] = lck

	return &lck, nil
}

func (l *localLockClient) KeepAlive(lock Lock) (Lock, error) {
	if len(lock.GetResource()) == 0 || lock.GetTTL() < 0 {
		return nil, ErrInvalidArg
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	lck, found := l.locks[lock.GetResource()]
	if !found {
		return nil, ErrNotFound
	}

	if lck.expired() {
		delete(l.locks, lock.GetResource())

		return nil, ErrExpired
	}

	if lck.GetToken() != lock.GetToken() {
		return nil, ErrInvalidArg
	}

	t := time.Now().Add(time.Duration(lock.GetTTL()) * time.Second)
	lck.TTL = lock.GetTTL()
	lck.expiration = &t

	return lock, nil
}

func (l *localLockClient) Unlock(lock Lock) error {
	if len(lock.GetResource()) == 0 || lock.GetTTL() < 0 {
		return ErrInvalidArg
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	lck, found := l.locks[lock.GetResource()]
	if !found {
		return ErrNotFound
	}

	if lck.expired() {
		delete(l.locks, lock.GetResource())

		return nil
	}

	if lck.GetToken() != lock.GetToken() {
		return ErrInvalidArg
	}

	delete(l.locks, lock.GetResource())

	return nil
}
