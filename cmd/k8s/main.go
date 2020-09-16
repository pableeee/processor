package main

import (
	"log"

	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

func main() {
	infra := infra.MakeNewInfra()

	err := infra.CreateServer("pable", "cs16")
	if err != nil {
		log.Fatalf("Could not create server: %s", err.Error())
	}

	log.Println("Game created")
}
