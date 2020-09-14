package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Owner     string    `json:"owner"`
	Game      string    `json:"game"`
	GameID    string    `json:"id"`
	CreatedAt time.Time `json:"created-at"`
}

type infraAPI interface {
	Get(userID string) ([]Server, error)
}

type LocalInfra struct {
	db  map[string][]Server
	mux *sync.Mutex
}

func (infra *LocalInfra) Get(userID string) ([]Server, error) {
	if len(userID) == 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	infra.mux.Lock()
	s, found := infra.db[userID]
	infra.mux.Unlock()

	if !found {
		return []Server{}, nil
	}

	return s, nil

}

func MakeLocalInfra() *LocalInfra {
	infra := new(LocalInfra)
	infra.mux = &sync.Mutex{}
	infra.db = make(map[string][]Server)
	return infra
}

func hangleGet(api infraAPI, w http.ResponseWriter, r *http.Request) error {
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

func main() {
	infra := MakeLocalInfra()
	infra.db["pableeee"] = []Server{
		{Game: "cs", CreatedAt: time.Now(), GameID: "1cc9b4f96ca4a8010d9575c4a121667c", Owner: "pableeee"},
	}

	infra.db["pecoreli"] = []Server{
		{Game: "l4d2", CreatedAt: time.Now(), GameID: "9a386f7947f5471c47eee8e18721054f", Owner: "pecoreli"},
		{Game: "cs", CreatedAt: time.Now(), GameID: "80380695fd1298fe2cd83f156e6c81c1", Owner: "pecoreli"},
	}

	r := mux.NewRouter()
	r.HandleFunc("/game/{userID}", func(w http.ResponseWriter, r *http.Request) {
		err := hangleGet(infra, w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	c := make(chan struct{})
	go func() {
		s := make(chan os.Signal)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)
		<-s
		srv.Close()
		c <- struct{}{}
	}()

	log.Fatal(srv.ListenAndServe())
	<-c
	os.Exit(0)
}
