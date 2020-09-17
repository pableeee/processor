package infra

import (
	"errors"
	"fmt"
)

var (
	ErrorUserNotFound = fmt.Errorf("user is missing")
	ErrorRetrieve     = fmt.Errorf("coulnd not retrieve")
	ErrorUpdate       = fmt.Errorf("coulnd not update")
	ErrorDelete       = fmt.Errorf("coulnd not delete")
	ErrorUnknown      = fmt.Errorf("unknown error")
)

// RequestHandler handler creation request
type RequestHandler struct {
	gameKVS *GameKVS
	userKVS *UserKVS
}

func (rh *RequestHandler) GetServer(userID string) ([]Server, error) {
	if len(userID) == 0 {
		return nil, ErrorUserNotFound
	}

	IDs, err := rh.userKVS.Get(userID)
	if err != nil || len(IDs) == 0 {
		return nil, ErrorRetrieve
	}

	srvs := make([]Server, len(IDs))

	for i, id := range IDs {
		s, err := rh.gameKVS.Get(id)
		if err != nil {
			return nil, ErrorRetrieve
		}

		srvs[i] = s
	}

	return srvs, nil
}

func (rh *RequestHandler) CreateServer(s Server) error {
	ids, err := rh.userKVS.Get(s.Owner)
	if errors.Is(err, UserNotFound) {
		ids = make([]string, 0)
	} else if err != nil {
		fmt.Printf("error: %s", err.Error())

		return ErrorRetrieve
	}

	ids = append(ids, s.GameID)

	err = rh.userKVS.Put(s.Owner, ids)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return ErrorUpdate
	}

	err = rh.gameKVS.Put(s.GameID, s)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return ErrorUpdate
	}

	return nil
}

func (rh *RequestHandler) DeleteServer(gameID string) error {
	s, err := rh.gameKVS.Get(gameID)
	if err != nil {
		return ErrorRetrieve
	}

	err = rh.gameKVS.Del(s.GameID)
	if err != nil {
		return ErrorDelete
	}

	ids, err := rh.userKVS.Get(s.Owner)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return ErrorUpdate
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

		return ErrorUpdate
	}

	return nil
}

func MakeRequestHandler(g *GameKVS, u *UserKVS) RequestHandler {
	rh := RequestHandler{gameKVS: g, userKVS: u}

	return rh
}
