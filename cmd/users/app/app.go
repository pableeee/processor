package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	ht "github.com/pableeee/processor/pkg/http"
	rep "github.com/pableeee/processor/pkg/repository"
)

type UserService struct {
	sv   *http.Server
	repo rep.Repository
}

type Profile struct {
	Type  string
	Group string
}

type User struct {
	ID       uint64
	Username string
	Mail     string
	Profile  Profile
}

func NewUserService() *UserService {
	s := &UserService{}
	s.sv = ht.DefaultBuilder().
		WithAddress("0.0.0.0").
		WithHandlerSetUp(
			func(r *mux.Router) {
				r.HandleFunc("/users/{ID}", func(w http.ResponseWriter, r *http.Request) {
					var usr User

					vars := mux.Vars(r)
					id := vars["kvsID"]

					if err := s.repo.Get(id, &usr); err != nil {
						w.WriteHeader(http.StatusInternalServerError)

						return
					}

					b, err := json.Marshal(usr)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)

						return
					}

					w.WriteHeader(http.StatusOK)
					w.Write(b)

				}).Methods("GET")
			},
		).
		Build()

	return s
}
