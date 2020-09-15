package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

type requestHandler struct {
	gameKVS *infra.GameKVS
	userKVS *infra.UserKVS
}

func (rh *requestHandler) handleUserGet(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if vars == nil {
		return fmt.Errorf("owner is missing")
	}

	own, found := vars["userID"]
	if !found {
		return fmt.Errorf("user id is missing")
	}

	IDs, err := rh.userKVS.Get(own)
	if err != nil || len(IDs) == 0 {
		return fmt.Errorf("coulnd not retrieve games")
	}

	var srvs []infra.Server

	for _, id := range IDs {
		s, err := rh.gameKVS.Get(id)
		if err != nil {
			return fmt.Errorf("coulnd not retrieve games")
		}
		srvs = append(srvs, s)
	}

	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s has %d active games\n", own, len(srvs)))

	for _, v := range srvs {
		b.WriteString(fmt.Sprintf("%s\n", v.Game))
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, b.String())

	return nil
}

func (rh *requestHandler) handleGamePost(w http.ResponseWriter, r *http.Request) error {

	d := json.NewDecoder(r.Body)
	s := infra.Server{}

	err := d.Decode(&s)
	s.GameID = uuid.New().String()

	if err != nil {
		log.Fatalf("error: %s", err.Error())

		return fmt.Errorf("could not decode mesage")
	}

	ids, err := rh.userKVS.Get(s.Owner)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	ids = append(ids, s.GameID)
	err = rh.userKVS.Put(s.Owner, ids)
	if err != nil {
		log.Fatalf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	err = rh.gameKVS.Put(s.Owner, s)
	if err != nil {
		log.Fatalf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "created")

	return nil
}

func (rh *requestHandler) handleGameDelete(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	if vars == nil {
		return fmt.Errorf("owner is missin")
	}

	gameID, found := vars["gameID"]
	if !found {
		return fmt.Errorf("user id is missing")
	}

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
		log.Fatalf("error: %s", err.Error())
		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	for i, id := range ids {
		if id == s.Owner {
			ids = append(ids[:i], ids[i+1:]...)
			break
		}
	}

	err = rh.userKVS.Put(s.Owner, ids)
	if err != nil {
		log.Fatalf("error: %s", err.Error())

		return fmt.Errorf("could not update servers for user:%s", s.Owner)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "created")

	return nil
}
