package jet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/k8s/builder"
	"github.com/pableeee/processor/pkg/k8s/provider"
	"github.com/pableeee/processor/pkg/k8s/provider/types"
	"github.com/pableeee/processor/pkg/kvs"
	"github.com/pableeee/processor/pkg/repository"
)

type Service struct {
	builder *builder.Builder
}

func NewJetService(kubeconfig string) *Service {
	s := Service{}
	p := provider.NewInfraProvider(kubeconfig)
	s.builder = builder.NewBuilder().
		WithProvider(p).
		WithRepository(repository.WithKVS(kvs.NewLocal()))

	return &s
}

func (j *Service) addHandlers(r *mux.Router) {
	// handles get games requests
	r.HandleFunc("/project/{projID}/kvs/{kvsID}", func(w http.ResponseWriter, r *http.Request) {
		kvsID := mux.Vars(r)["kvsID"]
		projID := mux.Vars(r)["projID"]

		s, err := j.builder.GetService(projID, kvsID, types.KVS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		b, err := json.Marshal(s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(b))

	}).Methods("GET")

	r.HandleFunc("/project/{projID}/kvs/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["projID"]

		if err := j.builder.
			BuildKVS(builder.Model{
				Project: id,
				Type:    types.KVS,
			}); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		if err := j.builder.BuildKVS(builder.Model{
			Project:      id,
			URL:          "",
			Repo:         "",
			ServivceName: "",
			Type:         types.KVS,
		}); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error: %s", err.Error())

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, map[string]interface{}{
			"status": 200,
			"msg":    fmt.Sprintf("KVS %s created in %s", id, "some name"),
		})

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
