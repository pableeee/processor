package lock

import "errors"

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
)
