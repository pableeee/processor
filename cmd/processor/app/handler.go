package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

var (
	ErrorUserMissing    = fmt.Errorf("owner is missing")
	ErrorRetrieve       = fmt.Errorf("error retieving the servers for user")
	ErrorDecodingMsg    = fmt.Errorf("could not decode mesage")
	ErrorCreatingServer = fmt.Errorf("could not create server")
	ErrorDeletingServer = fmt.Errorf("coulnd delete server")
)

type requestHandler struct {
	handler infra.InfraManager
}

func (rh *requestHandler) handleUserGet(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if vars == nil {
		return ErrorUserMissing
	}

	own, found := vars["userID"]
	if !found {
		return ErrorUserMissing
	}

	srvs, err := rh.handler.GetServer(own)
	if err != nil {
		return ErrorRetrieve
	}

	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s has %d active games\n", own, len(srvs)))

	for _, v := range srvs {
		b.WriteString(fmt.Sprintf("%s - id:%s\n", v.Game, v.GameID))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, b.String())

	return nil
}

func (rh *requestHandler) handleGamePost(w http.ResponseWriter, r *http.Request) error {
	d := json.NewDecoder(r.Body)
	s := infra.Server{}

	err := d.Decode(&s)
	if err != nil {
		fmt.Printf("error: %s", err.Error())

		return ErrorDecodingMsg
	}

	s, err = rh.handler.CreateServer(s.Owner, s.Game)
	if err != nil {
		return ErrorCreatingServer
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "game %#v created\n", s)

	return nil
}

func (rh *requestHandler) handleGameDelete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if vars == nil {
		return ErrorUserMissing
	}

	gameID, found := vars["gameID"]
	if !found {
		return ErrorUserMissing
	}

	err := rh.handler.DeleteServer(gameID)
	if err != nil {
		return ErrorDeletingServer
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "created")

	return nil
}
