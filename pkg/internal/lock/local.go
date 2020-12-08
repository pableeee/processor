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

type lock struct {
	TTL        int
	Resource   string
	Origin     string
	Token      string
	expiration *time.Time
}

func (l *lock) expired() bool {
	if l.TTL == 0 || l.expiration == nil {
		return false
	}

	return l.expiration.Before(time.Now())
}

func (l *lock) GetResource() string {
	return l.Resource
}

func (l *lock) SetResource(resource string) {
	l.Resource = resource
}

func (l *lock) GetTTL() int {
	return l.TTL
}

func (l *lock) SetTTL(TTL int) {
	l.TTL = TTL
}

func (l *lock) GetOrigin() string {
	return l.Origin
}

func (l *lock) SetOrigin(origin string) {
	l.Origin = origin
}

func (l *lock) GetToken() string {
	return l.Token
}

func (l *lock) SetToken(token string) {
	l.Token = token
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

func (l *localLockClient) Lock(resource string, TTL int) (Lock, error) {
	if len(resource) == 0 || TTL < 0 {
		return nil, ErrInvalidArg
	}

	l.mux.Lock()
	defer l.mux.Unlock()

	lck, found := l.locks[resource]
	if found && !lck.expired() {
		return nil, ErrLocked
	}

	token := uuid.New()
	t := time.Now().Add(time.Duration(TTL) * time.Second)
	lck = lock{
		TTL:        TTL,
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
