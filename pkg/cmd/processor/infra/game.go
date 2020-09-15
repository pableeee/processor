package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/pableeee/processor/pkg/internal/kvs"
)

var instance *GameKVS

// GameKVS for local tests
type GameKVS struct {
	kvs kvs.KVS
	mux *sync.Mutex
}

// Get implements a local get service
func (infra *GameKVS) Get(gameID string) (Server, error) {
	if len(gameID) == 0 {
		return Server{}, fmt.Errorf("invalid game id")
	}

	infra.mux.Lock()
	b, err := infra.kvs.Get(gameID)
	infra.mux.Unlock()

	if err != nil {
		log.Fatalf("could not retieve game from local repository: %s", err.Error())
		return Server{}, err
	}

	var s Server
	err = json.Unmarshal(b, &s)
	if err != nil {
		log.Fatalf("unable to unmarshall object from kvs: %s", err.Error())
		return Server{}, err
	}

	return s, nil

}

// Put implements a o local put service
func (infra *GameKVS) Put(gameID string, s Server) error {
	if len(gameID) == 0 {
		return fmt.Errorf("invalid game id")
	}

	b, err := json.Marshal(&s)
	if err != nil {
		log.Fatalf("error mashalling the game: %s", err.Error())

		return err
	}

	infra.mux.Lock()
	err = infra.kvs.Put(gameID, b)
	infra.mux.Unlock()

	if err != nil {
		log.Fatalf("error in kvs put for game: %s", err.Error())

		return err
	}

	return nil
}

// Get implements a local get service
func (infra *GameKVS) Del(gameID string) error {
	if len(gameID) == 0 {
		return fmt.Errorf("invalid game id")
	}

	infra.mux.Lock()
	defer infra.mux.Unlock()

	err := infra.kvs.Del(gameID)
	if err != nil {
		log.Fatalf("could not delete game from repository: %s", err.Error())
		return err
	}

	return nil
}

func MakeLocalGameRepository() *GameKVS {
	return makeGameKVS(kvs.MakeLocalKVS())
}

// MakeGameKVS makes an instances of a local infra
func makeGameKVS(kvs kvs.KVS) *GameKVS {
	if instance == nil {
		instance = new(GameKVS)
		instance.mux = &sync.Mutex{}
		instance.kvs = kvs
	}

	return instance
}
