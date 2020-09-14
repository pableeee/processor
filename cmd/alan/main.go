package main

import (
	"fmt"
	"os"

	"github.com/pableeee/processor/pkg/cmd/alan/get"
	"github.com/pableeee/processor/pkg/cmd/alan/put"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: "alan",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(get.NewCommand())
	rootCmd.AddCommand(put.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

}
