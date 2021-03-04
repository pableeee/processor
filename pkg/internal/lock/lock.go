package lock

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Lock interface {
	GetResource() string
	SetResource(string)
	GetTTL() int
	SetTTL(int)
	GetOrigin() string
	SetOrigin(string)
	GetToken() string
	SetToken(string)
}

type Client interface {
	Get(resource string) (Lock, error)
	Lock(resource string, ttl int) (Lock, error)
	KeepAlive(lock Lock) (Lock, error)
	Unlock(lock Lock) error
}

var (
	ErrNotFound   = errors.New("lock not found")
	ErrLocked     = errors.New("resource locked")
	ErrInvalidArg = errors.New("invalid arg")
	ErrExpired    = errors.New("lock expired")
	ErrInternal   = errors.New("internal error")
)

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

func (l *lock) SetTTL(ttl int) {
	l.TTL = ttl
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

func NewLock(resource string, ttl int) Lock {
	token := uuid.New()
	t := time.Now().Add(time.Duration(ttl) * time.Second)
	lck := lock{
		TTL:        ttl,
		Resource:   resource,
		Origin:     "",
		Token:      token.String(),
		expiration: &t,
	}

	return &lck
}
