package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func Run() error {
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

}
