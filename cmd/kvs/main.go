package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/pableeee/processor/pkg/cmd/kvs"
	"github.com/spf13/cobra"
)

var ErrInvalidParams = errors.New("invalid number of params")

func newLocalKVS() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(0),
		Use:     "local",
		Short:   "creates local kvs instace",
		Long:    "creates local kvs instace",
		Example: "kvs create local",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return ErrInvalidParams
			}

			s, err := kvs.NewLocalKVS()
			if err != nil {
				return err
			}

			s.Listen()

			return nil
		},
	}

	return cmd
}

func getCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "get",
		Short:   "gets key from kvs",
		Long:    "gets key from kvs",
		Example: "kvs get key",
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			c, err := kvs.NewKVSClient("localhost", 33333)
			if err != nil {
				return err
			}

			b, err := c.Get(key)
			if err != nil {
				return err
			}

			fmt.Printf("%s: %s", key, string(b))

			return nil
		},
	}

	return cmd
}

func putCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(2),
		Use:     "put",
		Short:   "puts key from kvs",
		Long:    "puts key from kvs",
		Example: "kvs put key value",
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]

			c, err := kvs.NewKVSClient("localhost", 33333)
			if err != nil {
				return err
			}

			b, err := json.Marshal(value)
			if err != nil {
				return err
			}

			err = c.Put(key, b)
			if err != nil {
				return err
			}

			fmt.Printf("%s: %s", key, string(b))

			return nil
		},
	}

	return cmd
}

func newRedisKVS() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "redis",
		Short:   "creates redis kvs instace",
		Long:    "creates redis kvs instace",
		Example: "kvs create redis port (localhost)",
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			s, err := kvs.NewRedisKVS(p)
			if err != nil {
				return err
			}

			s.Listen()

			return nil
		},
	}

	return cmd
}

func main() {
	rootCmd := &cobra.Command{
		Use: "kvs",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getCommand())
	rootCmd.AddCommand(putCommand())

	createCmd.AddCommand(newLocalKVS())
	createCmd.AddCommand(newRedisKVS())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
