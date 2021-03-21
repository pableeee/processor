package svc

import (
	"fmt"

	"github.com/pableeee/processor/pkg/internal/k8s"
	"github.com/pableeee/processor/pkg/k8s/builder"
	"github.com/pableeee/processor/pkg/k8s/provider"
	"github.com/pableeee/processor/pkg/k8s/provider/types"
	"github.com/pableeee/processor/pkg/kvs"
	"github.com/pableeee/processor/pkg/lock"
	"github.com/spf13/cobra"
)

var (
	ErrorInvalidArgs = fmt.Errorf("invalid arguments")
	ErrorMarshalling = fmt.Errorf("unable to marshal info to backend")
)

// NewCommand returns a new cobra.Command for cluster creation.
func NewCommand() *cobra.Command {
	root := newCommand("service", nil)
	create := newCommand("create", nil)

	create.AddCommand(buildLockCommand())
	create.AddCommand(buildKVSCommand())
	create.AddCommand(buildQueueCommand())
	root.AddCommand(create)

	return root
}

func buildLockCommand() *cobra.Command {
	cmd := buildCommand("lock", "alan service create lock MY_APP LOCK_NAME",
		func(b *builder.Builder, m builder.Model) error {
			return b.BuildLock(m)
		})

	return cmd
}

func buildKVSCommand() *cobra.Command {
	cmd := buildCommand("kvs", "alan service create kvs MY_APP KVS_NAME",
		func(b *builder.Builder, m builder.Model) error {
			return b.BuildKVS(m)
		})

	return cmd
}

func buildQueueCommand() *cobra.Command {
	cmd := buildCommand("queue", "alan service create queue MY_APP QUEUE_NAME",
		func(b *builder.Builder, m builder.Model) error {
			return b.BuildQueue(m)
		})

	return cmd
}

func buildCommand(use, example string, f func(b *builder.Builder, m builder.Model) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Example: example,
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			repo := args[1]
			url := args[2]
			service := args[3]

			kubeconfig, err := k8s.GetKubeConfig()
			if err != nil {
				return fmt.Errorf("failed getting kubeconfig path: %w", err)
			}

			p := provider.NewInfraProvider(kubeconfig)
			b := builder.NewBuilder().
				WithProvider(p).
				WithLock(lock.NewLocal()).
				WithKVS(kvs.NewLocal())

			if err = f(b, builder.Model{
				Project:      name,
				Repo:         repo,
				URL:          url,
				ServivceName: service,
				Type:         types.Lock,
			}); err != nil {
				return fmt.Errorf("failed building lock service: %w", err)
			}

			return nil
		},
	}

	return cmd
}

func newCommand(use string, f func(cmd *cobra.Command, args []string) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:  use,
		RunE: f,
	}

	return cmd
}
