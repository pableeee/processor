package jet

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/internal/kvs"
	"github.com/pableeee/processor/pkg/internal/lock"
	"github.com/pableeee/processor/pkg/k8s/builder"
	"github.com/pableeee/processor/pkg/k8s/provider"
)

type Service struct {
	builder *builder.Builder
}

func NewJetService(kubeconfig string) *Service {
	s := Service{}
	p := provider.NewInfraProvider(kubeconfig)
	s.builder = builder.NewBuilder().
		WithProvider(p).
		WithLock(lock.NewLocal()).
		WithKVS(kvs.NewLocal())

	return &s
}

func (j *Service) addHandlers(r *mux.Router) {
	// handles get games requests
	r.HandleFunc("/kvs/{kvsID}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["kvsID"]

		svcs := builder.NewBuilder().GetServices()

		b, err := svcs.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(b))

	}).Methods("GET")

	r.HandleFunc("/kvs/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["kvsID"]
		svcs := builder.NewBuilder().GetServices()

		b, err := svcs.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(b))

	}).Methods("POST")

	// handles get games requests
	r.HandleFunc("/game/{gameID}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", "error")
	}).Methods("DELETE")

	// handles post to add new games
	r.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", "")
	}).Methods("POST")
}
