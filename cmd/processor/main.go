package main

import (
	"log"
	"os"

	"github.com/pableeee/processor/cmd/processor/app"
)

func main() {
	/*
		i := infra.MakeLocalInfra()
		i.Put("pableeee", infra.Server{
			Game: "cs", CreatedAt: time.Now(),
			GameID: "1cc9b4f96ca4a8010d9575c4a121667c",
			Owner:  "pableeee"})

		i.Put("pecoreli", infra.Server{Game: "l4d2", CreatedAt: time.Now(), GameID: "9a386f7947f5471c47eee8e18721054f", Owner: "pecoreli"})
		i.Put("pecoreli", infra.Server{Game: "cs", CreatedAt: time.Now(), GameID: "80380695fd1298fe2cd83f156e6c81c1", Owner: "pecoreli"})
	*/
	if err := app.Main(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
