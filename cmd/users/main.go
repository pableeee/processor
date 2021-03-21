package main

import (
	"log"

	"github.com/pableeee/processor/cmd/users/app"
)

func main() {
	us := app.NewUserService(nil)
	if err := us.Start(); err != nil {
		log.Fatal("could not start server")
	}

}
