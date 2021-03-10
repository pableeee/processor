package main

import (
	"log"
	"os"

	"github.com/pableeee/processor/pkg/cmd/alan/del"
	"github.com/pableeee/processor/pkg/cmd/alan/get"
	"github.com/pableeee/processor/pkg/cmd/alan/put"
	"github.com/pableeee/processor/pkg/cmd/alan/svc"
	"github.com/pableeee/processor/pkg/cmd/alan/write"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "alan",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(get.NewCommand())
	rootCmd.AddCommand(put.NewCommand())
	rootCmd.AddCommand(del.NewCommand())
	rootCmd.AddCommand(write.NewCommand())
	rootCmd.AddCommand(svc.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

}
