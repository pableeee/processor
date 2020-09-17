package main

import (
	"fmt"
	"log"

	"github.com/pableeee/processor/pkg/cmd/processor/infra"
)

func main() {
	infra := infra.MakeNewInfra()

	s, err := infra.CreateServer("pable", "cs16")
	if err != nil {
		log.Fatalf("Could not create server: %s", err.Error())
	}

	fmt.Printf("Game created!\n%#v", s)
}
