package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pableeee/processor/pkg/internal/k8s"
)

type protocol uint16

const (
	Invalid protocol = iota
	TCP
	UDP
)

// Server represents a game server
type Server struct {
	Owner     string         `json:"owner"`
	Game      string         `json:"game"`
	GameID    string         `json:"id"`
	Ports     []PortSettings `json:"ports"`
	CreatedAt time.Time      `json:"created-at"`
	IP        string         `json:"ip"`
}

type PortSettings struct {
	Proto    protocol `json:"protocol"`
	NodePort int64    `json:"node-port"`
	Port     int64    `json:"port"`
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

func (i *Infra) CreateServer(userID, game string) (Server, error) {
	podID := "p" + uuid.New().String()[:8]
	s := Server{Game: game, CreatedAt: time.Now(), Owner: userID, GameID: podID}
	img := i.mapper.GetImage(game)
	// TODO: setting +1 replicas of an existing deployment if possible
	_, err := i.deploy.CreateDeployment("", "default", img, podID)
	if err != nil {
		return s, fmt.Errorf("error creating resource: %s", err.Error())
	}
	time.Sleep(1 * time.Second)
	//CreateService(cfg, namespace, name string, port uint16) (ServiceResponse, error)
	res, err := i.svc.CreateService("", "default", podID, 80)
	if err != nil {
		//TODO: pods was created, but not the service
		return s, fmt.Errorf("error creating resource: %s", err.Error())
	}

	ports, err := res.GetSlice("spec", "ports")
	if err != nil {
		//TODO: pods was created, but not the service
		return s, fmt.Errorf("error creating resource: %s", err.Error())
	}

	s.IP, _ = res.GetString("spec", "clusterIP")
	s.Ports = make([]PortSettings, len(ports))
	for i, v := range ports {
		m, ok := v.(map[string]interface{})
		if !ok {
			return s, fmt.Errorf("Unable to unwrap service port details")
		}
		p := PortSettings{}
		proto := m["protocol"]
		switch proto {
		case "TCP":
			p.Proto = TCP
		case "UDP":
			p.Proto = UDP
		}

		p.NodePort = m["nodePort"].(int64)
		p.Port = m["port"].(int64)
		s.Ports[i] = p
	}

	return s, nil
}

func (i *Infra) DeleteServer(gameID string) error {

	return nil
}
