package infra

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/pableeee/processor/pkg/kvs"
)

var (
	ErrorGameNotFound = fmt.Errorf("game not found")
	ErrorInvalidID    = fmt.Errorf("invalid id")
)

// GameKVS for local tests
type GameKVS struct {
	kvs kvs.KVS
}

// Get implements a local get service
func (infra *GameKVS) Get(gameID string) (Server, error) {
	if len(gameID) == 0 {
		return Server{}, ErrorInvalidID
	}

	b, err := infra.kvs.Get(gameID)
	if errors.Is(err, kvs.ErrKeyNotFound) {
		return Server{}, nil
	} else if err != nil {
		fmt.Printf("could not retieve game from local repository: %s", err.Error())

		return Server{}, err
	}

	var s Server

	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("unable to unmarshall object from kvs: %s", err.Error())

		return Server{}, err
	}

	return s, nil
}

// Put implements a o local put service
func (infra *GameKVS) Put(gameID string, s Server) error {
	if len(gameID) == 0 {
		return ErrorInvalidID
	}

	b, err := json.Marshal(&s)
	if err != nil {
		fmt.Printf("error mashalling the game: %s", err.Error())

		return err
	}

	err = infra.kvs.Put(gameID, b)
	if err != nil {
		fmt.Printf("error in kvs put for game: %s", err.Error())

		return err
	}

	return nil
}

// Get implements a local get service
func (infra *GameKVS) Del(gameID string) error {
	if len(gameID) == 0 {
		return ErrorInvalidID
	}

	err := infra.kvs.Del(gameID)
	if err != nil {
		fmt.Printf("could not delete game from repository: %s", err.Error())

		return err
	}

	return nil
}

func MakeLocalGameRepository() *GameKVS {
	return makeGameKVS(kvs.MakeLocalKVS())
}

// MakeGameKVS makes an instances of a local infra
func makeGameKVS(kvs kvs.KVS) *GameKVS {
	instance := new(GameKVS)
	instance.kvs = kvs

	return instance
}
