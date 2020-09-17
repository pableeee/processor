package infra

import "fmt"

type InfraManager interface {
	CreateServer(userID, game string) (Server, error)
	DeleteServer(gameID string) error
	GetServer(userID string) ([]Server, error)
}

type InfraService struct {
	mng *Infra
	rh  RequestHandler
}

func (is *InfraService) CreateServer(userID, game string) (Server, error) {
	if len(userID) == 0 || len(game) == 0 {
		return Server{}, fmt.Errorf("invalid arguments")
	}

	s, err := is.mng.CreateServer(userID, game)
	if err != nil {
		return Server{}, err
	}

	err = is.rh.CreateServer(s)
	if err != nil {
		return Server{}, err
	}

	return s, nil
}

func (is *InfraService) DeleteServer(gameID string) error {
	if len(gameID) == 0 {
		return fmt.Errorf("invalid arguments")
	}

	err := is.mng.DeleteServer(gameID)
	if err != nil {
		return err
	}

	err = is.rh.DeleteServer(gameID)
	if err != nil {
		return err
	}

	return nil
}

func (is *InfraService) GetServer(userID string) ([]Server, error) {
	return is.rh.GetServer(userID)
}

func MakeInfraService() InfraManager {
	h := MakeRequestHandler(MakeLocalGameRepository(),
		MakeLocalUserRepository())
	m := MakeNewInfra()
	i := &InfraService{rh: h, mng: m}
	return i
}
