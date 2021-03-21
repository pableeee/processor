package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/pableeee/processor/pkg/internal/kvs"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

// The repository packaget, provides the basis of an entity repository, to avoid boilerplate code.

// Repository abstracts the actual underlyin db infra, into a client interface.
type Repository interface {
	Get(id string, i interface{}) error
	Save(id string, i interface{}) error
	Update(id string, i interface{}) error
	Delete(id string) error
}

type repository struct {
	store kvs.KVS
}

func WithKVS(store kvs.KVS) Repository {
	return &repository{store: store}
}

func (r repository) Delete(id string) error {
	if len(id) == 0 {
		return ErrInvalidKey
	}

	if err := r.store.Del(id); err != nil {
		return fmt.Errorf("unable to delete %s: %w", id, err)
	}

	return nil
}

func (r repository) Get(id string, i interface{}) error {
	if len(id) == 0 {
		return ErrInvalidKey
	}

	b, err := r.store.Get(id)
	if err != nil {
		return fmt.Errorf("fail retrieving id:%s %w", id, err)
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return fmt.Errorf("fail unmarshaling response :%s %w", string(b), err)
	}

	return nil
}

func (r repository) Save(id string, i interface{}) error {
	if len(id) == 0 {
		return ErrInvalidKey
	}

	b, err := json.Marshal(i)
	if err != nil {
		// TODO cambiar error
		return ErrInvalidKey
	}

	if err = r.store.Put(id, b); err != nil {
		return fmt.Errorf("failed saving %s: %w", id, err)
	}

	return nil
}

func (r repository) Update(id string, i interface{}) error {
	if len(id) == 0 {
		return ErrInvalidKey
	}

	prototype := reflect.TypeOf(i)
	if prototype.Kind() == reflect.Ptr {
		prototype = prototype.Elem()
	}

	instance := reflect.New(prototype).Interface()

	b, err := r.store.Get(id)
	if err != nil {
		return fmt.Errorf("failed retrieving %s: %w", id, err)
	}

	if err = json.Unmarshal(b, &instance); err != nil {
		return fmt.Errorf("failed unmashaling previous value %s: %w", id, err)
	}

	if err = r.merge(instance, i); err != nil {
		return fmt.Errorf("failed merging values %s: %w", id, err)
	}

	if b, err = json.Marshal(instance); err != nil {
		return fmt.Errorf("failed mashaling values %s: %w", id, err)
	}

	if err = r.store.Put(id, b); err != nil {
		return fmt.Errorf("failed updating values %s: %w", id, err)
	}

	return nil
}

func (r repository) merge(dst, org interface{}) error {
	if reflect.TypeOf(dst) != reflect.TypeOf(org) {
		return fmt.Errorf("non compatible types")
	}

	return nil
}
