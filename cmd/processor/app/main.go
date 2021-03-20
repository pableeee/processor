package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

var handler requestHandler

func addHandlers(r *mux.Router) {
	// handles get games requests
	r.HandleFunc(fmt.Sprintf("/project/{%s}", proyectID), func(w http.ResponseWriter, r *http.Request) {
		err := handler.handleProyectGet(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/project/{%s}", proyectID), func(w http.ResponseWriter, r *http.Request) {
		err := handler.handleProyectGet(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}).Methods("POST")

	// handles get games requests
	r.HandleFunc("/game/{gameID}", func(w http.ResponseWriter, r *http.Request) {
		err := handler.handleGameDelete(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}).Methods("DELETE")

	// handles post to add new games
	r.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		err := handler.handleGamePost(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
	}).Methods("POST")
}

func makeServer() *http.Server {
	r := mux.NewRouter()
	addHandlers(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}

func handleSignals() <-chan struct{} {
	c := make(chan struct{})

	handleKill := func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)
		// waiting for SIGTERM, to stop de server
		<-s
		// signal to exit the app
		c <- struct{}{}
	}

	go handleKill()

	return c
}

// Run executes the main app loop.
func Run() error {
	srv := makeServer()
	handler.handler = infra.MakeInfraService()
	c := handleSignals()
	// start listening
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	// waiting for server to shutdown
	<-c

	log.Println("shuting down server")
	srv.Close()

	return nil
}

// Main app function.
func Main() error {
	return Run()
}
