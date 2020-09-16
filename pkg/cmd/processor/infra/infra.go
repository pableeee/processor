package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pableeee/processor/pkg/cmd/k8s"
)

type protocol uint16

const (
	Invalid protocol = iota
	TCP
	UDP
)

// Server represents a game server
type Server struct {
	Owner     string              `json:"owner"`
	Game      string              `json:"game"`
	GameID    string              `json:"id"`
	Ports     map[protocol]uint16 `json:"ports"`
	CreatedAt time.Time           `json:"created-at"`
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
	deploy k8s.DeploymentManager
	svc    k8s.ServiceManager
	mapper ImageMapper
}

func MakeNewInfra() *Infra {
	i := new(Infra)
	i.deploy = &k8s.DeploymentManagerImpl{}
	i.svc = &k8s.ServiceManagerImpl{}
	i.mapper = &trivialMapper{}
	return i
}

type ImageMapper interface {
	GetImage(game string) string
}
type trivialMapper struct {
}

func (t *trivialMapper) GetImage(game string) string {
	return "nginx"
}

func (i *Infra) CreateServer(userID, game string) error {
	podID := uuid.New().String()[:8]
	img := i.mapper.GetImage(game)
	// TODO: setting +1 replicas of an existing deployment if possible
	_, err := i.deploy.CreateDeployment("", "default", img, podID)
	if err != nil {
		return fmt.Errorf("error creating resource: %s", err.Error())
	}

	//CreateService(cfg, namespace, name string, port uint16) (ServiceResponse, error)
	_, err = i.svc.CreateService("", "default", podID, 80)
	if err != nil {
		//TODO: pods was created, but not the service
		return fmt.Errorf("error creating resource: %s", err.Error())
	}

	//res.
	return nil
}

func (i *Infra) DeleteServer(gameID string) error {

	return nil
}
