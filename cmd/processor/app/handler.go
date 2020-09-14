package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

func hangleGet(api infra.Backend, w http.ResponseWriter, r *http.Request) error {
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
