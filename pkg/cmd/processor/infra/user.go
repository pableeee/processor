package infra

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/pableeee/processor/pkg/kvs"
)

// UserKVS for local tests
type UserKVS struct {
	kvs kvs.KVS
}

var (
	UserNotFound = errors.New("User not found")
)

// Get implements a local get service
func (infra *UserKVS) Get(userID string) ([]string, error) {
	if len(userID) == 0 {
		return []string{}, fmt.Errorf("invalid game id")
	}

	b, err := infra.kvs.Get(userID)

	if err == kvs.ErrKeyNotFound {
		return []string{}, UserNotFound
	} else if err != nil {
		fmt.Printf("could not retieve game from local repository: %s", err.Error())
		return []string{}, err
	}

	var s []string
	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("unable to unmarshall object from kvs: %s", err.Error())
		return []string{}, err
	}

	return s, nil

}

// Put implements a o local put service
func (infra *UserKVS) Put(userID string, IDs []string) error {
	if len(userID) == 0 {
		return fmt.Errorf("invalid game id")
	}

	b, err := json.Marshal(&IDs)
	if err != nil {
		fmt.Printf("error mashalling the games: %s", err.Error())

		return err
	}

	err = infra.kvs.Put(userID, b)
	if err != nil {
		fmt.Printf("error in kvs put for game: %s", err.Error())

		return err
	}

	return nil
}

// Get implements a local get service
func (infra *UserKVS) Del(userID string) error {
	if len(userID) == 0 {
		return fmt.Errorf("invalid game id")
	}

	err := infra.kvs.Del(userID)
	if err != nil {
		fmt.Printf("could not delete games from repository: %s", err.Error())
		return err
	}

	return nil
}

func MakeLocalUserRepository() *UserKVS {
	return makeUserKVS(kvs.MakeLocalKVS())
}

// MakeUserKVS makes an userInstances of a local infra
func makeUserKVS(kvs kvs.KVS) *UserKVS {
	userInstance := new(UserKVS)
	userInstance.kvs = kvs
	return userInstance
}
