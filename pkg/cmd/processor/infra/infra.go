package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pableeee/processor/pkg/cmd/k8s"
)

// Server represents a game server
type Server struct {
	Owner     string    `json:"owner"`
	Game      string    `json:"game"`
	GameID    string    `json:"id"`
	CreatedAt time.Time `json:"created-at"`
}

// Backend represents the backend server storing
type UserRepository interface {
	Get(userID string) ([]string, error)
	Put(userID string, s Server) error
	Del(userID, gameID string) error
}

type GameRepository interface {
	Get(gameID string) (Server, error)
	Put(gameID string, s Server) error
	Del(gameID string) error
}

type Infra struct {
	kvs UserRepository
	dm  k8s.DeploymentManager
}

func (i *Infra) CreateService(userID, game string) error {
	uuid := uuid.New()
	_, err := i.dm.CreateDeployment("", "default", game, uuid.String())
	if err != nil {
		return fmt.Errorf("error creating resource: %s", err.Error())
	}

	s := Server{Game: game, CreatedAt: time.Now(), GameID: uuid.String(), Owner: userID}
	k := fmt.Sprintf("%s:%s", s.Owner, s.GameID)

	err = i.kvs.Put(k, s)
	if err != nil {
		return fmt.Errorf("error on kvs put: %s", err.Error())
	}

	return nil
}

func (i *Infra) DeleteService(gameID string) error {
	i.kvs.Get(gameID)
	return nil
}
