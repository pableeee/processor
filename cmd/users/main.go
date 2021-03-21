package main

import (
	"log"

	"github.com/pableeee/processor/cmd/users/app"
	"github.com/pableeee/processor/pkg/kvs"
)

func main() {
	us := app.NewUserService(kvs.NewLocal())
	if err := us.Start(); err != nil {
		log.Fatal("could not start server")
	}

}
