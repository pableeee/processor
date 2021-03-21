package app

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	ht "github.com/pableeee/processor/pkg/http"
	"github.com/pableeee/processor/pkg/kvs"
	"github.com/pableeee/processor/pkg/repository"
	rep "github.com/pableeee/processor/pkg/repository"
)

type UserService struct {
	sv   *http.Server
	repo rep.Repository
}

type Profile struct {
	Type  string `json:"type,omitempty"`
	Group string `json:"group,omitempty"`
}

type User struct {
	ID       uint64  `json:"id,omitempty"`
	Username string  `json:"username,omitempty"`
	Mail     string  `json:"mail,omitempty"`
	Profile  Profile `json:"profile,omitempty"`
}

func (s *UserService) Start() error {
	// Wait for the user to cancel the process
	if s.sv == nil {
		return fmt.Errorf("server not configured")
	}

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		// Close server on SIGINT
		ctx := context.Background()
		s.sv.Shutdown(ctx)
	}()

	if err := s.sv.ListenAndServe(); err != nil {
		return fmt.Errorf("faild starting server %w", err)
	}

	return nil
}

func handleGet(s *UserService, r *http.Request) ([]byte, error) {
	var usr User

	vars := mux.Vars(r)
	id := vars["ID"]

	if err := s.repo.Get(id, &usr); err != nil {
		return nil, fmt.Errorf("failed getting user %s: %w", id, err)
	}

	b, err := json.Marshal(usr)
	if err != nil {
		return nil, fmt.Errorf("failed extracting object user %s: %w", id, err)
	}

	return b, nil
}

func handlePost(s *UserService, r *http.Request) ([]byte, error) {
	var usr User

	input := make([]byte, 0)
	scan := bufio.NewScanner(r.Body)

	for scan.Scan() {
		b := scan.Bytes()
		input = append(input, b...)
	}

	err := scan.Err()
	if scan.Err() != nil {
		return nil, fmt.Errorf("failed reading body: %w", err)
	}

	if err = json.Unmarshal(input, &usr); err != nil {
		return nil, fmt.Errorf("failed mashaling body: %w", err)
	}

	if err := s.repo.Save(fmt.Sprintf("%d", usr.ID), usr); err != nil {
		return nil, fmt.Errorf("failed saving user %w", err)
	}

	return nil, nil
}

func NewUserService(k kvs.KVS) *UserService {
	s := &UserService{}
	s.repo = repository.WithKVS(k)
	s.sv = ht.DefaultBuilder().
		WithAddress("0.0.0.0").
		WithPort(8888).
		WithHandlerSetUp(
			func(r *mux.Router) {
				// Handles GET
				r.HandleFunc("/users/{ID}", func(w http.ResponseWriter, r *http.Request) {
					b, err := handleGet(s, r)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)

						return
					}

					w.WriteHeader(http.StatusOK)
					// Hacer algun retry?
					_, _ = w.Write(b)
				}).Methods("GET")

				r.HandleFunc("/users/{ID}", func(w http.ResponseWriter, r *http.Request) {
					b, err := handlePost(s, r)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)

						return
					}

					w.WriteHeader(http.StatusOK)
					// Hacer algun retry?
					_, _ = w.Write(b)
				}).Methods("PUT")

				r.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
					_, err := handlePost(s, r)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(err.Error()))

						return
					}

					w.WriteHeader(http.StatusOK)
				}).Methods("POST")
			},
		).
		Build()

	return s
}
