package infra

import (
	"fmt"
)

// RequestHandler handler creation request
type RequestHandler struct {
	gameKVS *GameKVS
	userKVS *UserKVS
}

func (rh *RequestHandler) GetServer(userID string) ([]Server, error) {
	if len(userID) == 0 {
		return nil, fmt.Errorf("user is missing")
	}

	IDs, err := rh.userKVS.Get(userID)
	if err != nil || len(IDs) == 0 {
		return nil, fmt.Errorf("coulnd not retrieve games")
	}

	var srvs []Server

	for _, id := range IDs {
		s, err := rh.gameKVS.Get(id)
		if err != nil {
			return nil, fmt.Errorf("coulnd not retrieve games")
		}
		srvs = append(srvs, s)
	}

	return srvs, nil
}

func (rh *RequestHandler) CreateServer(s Server) error {
	ids, err := rh.userKVS.Get(s.Owner)
	if err == UserNotFound {
		ids = make([]string, 0)
	} else if err != nil {
		fmt.Printf("error: %s", err.Error())
		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	ids = append(ids, s.GameID)
	err = rh.userKVS.Put(s.Owner, ids)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	err = rh.gameKVS.Put(s.GameID, s)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	return nil
}

func (rh *RequestHandler) DeleteServer(gameID string) error {
	s, err := rh.gameKVS.Get(gameID)
	if err != nil {
		return fmt.Errorf("coulnd get servers for get kvs: %s", gameID)
	}

	err = rh.gameKVS.Del(s.GameID)
	if err != nil {
		return fmt.Errorf("coulnd delete servers for game kvs: %s", gameID)
	}

	ids, err := rh.userKVS.Get(s.Owner)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	for i, id := range ids {
		if id == s.GameID {
			ids = append(ids[:i], ids[i+1:]...)
			break
		}
	}

	err = rh.userKVS.Put(s.Owner, ids)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	return nil
}

func MakeRequestHandler(g *GameKVS, u *UserKVS) RequestHandler {
	rh := RequestHandler{gameKVS: g, userKVS: u}
	return rh
}
