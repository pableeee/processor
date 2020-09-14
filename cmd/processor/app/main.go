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

func getInfra() infra.Backend {
	return infra.MakeLocalInfra()
}

func makeServer() *http.Server {
	i := getInfra()
	r := mux.NewRouter()

	r.HandleFunc("/game/{userID}", func(w http.ResponseWriter, r *http.Request) {
		err := hangleGet(i, w, r)
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

	return srv
}

func handleSignals() <-chan struct{} {
	c := make(chan struct{})
	go func() {
		s := make(chan os.Signal)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)
		// waiting for SIGTERM, to stop de server
		<-s
		// signal to exit the app
		c <- struct{}{}
	}()
	return c
}

// Run executes the main app loop
func Run() error {
	srv := makeServer()

	// start listening
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	c := handleSignals()
	// waiting for server to shutdown
	<-c

	log.Println("shuting down server")
	srv.Close()

	return nil
}

// Main app function
func Main() error {
	return Run()
}
