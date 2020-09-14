package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

func handleGet(api infra.Repository, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if vars == nil {
		return fmt.Errorf("owner is missin")
	}

	own, found := vars["userID"]
	if !found {
		return fmt.Errorf("user id is missing")
	}

	srvs, err := api.Get(own)
	if err != nil || len(srvs) == 0 {
		return fmt.Errorf("coulnd not retrieve games")
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

func handlePost(api infra.Repository, w http.ResponseWriter, r *http.Request) error {

	d := json.NewDecoder(r.Body)
	s := infra.Server{}

	err := d.Decode(&s)

	if err != nil {
		return fmt.Errorf("could not decode mesage")
	}

	err = api.Put(s.Owner, s)
	if err != nil {
		return fmt.Errorf("coulnd update servers for user:%s", s.Owner)
	}

	return nil
}
