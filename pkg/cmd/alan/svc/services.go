package svc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	guuid "github.com/google/uuid"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
	"github.com/pableeee/processor/pkg/internal/k8s"
	"github.com/pableeee/processor/pkg/k8s/builder"
	"github.com/pableeee/processor/pkg/k8s/provider"
	"github.com/pableeee/processor/pkg/k8s/provider/types"
	"github.com/spf13/cobra"
)

const (
	url         = "http://127.0.0.1:8000/game/"
	contentType = "application/json"
)

var (
	ErrorInvalidArgs = fmt.Errorf("invalid arguments")
	ErrorMarshalling = fmt.Errorf("unable to marshal info to backend")
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	root := newCommand("service", nil)
	create := newCommand("create", nil)

	create.AddCommand(newCommand("lock", func(cmd *cobra.Command, args []string) error {
		name := args[0]
		repo := args[1]
		url := args[2]
		service := args[3]

		kubeconfig, err := k8s.GetKubeConfig()
		if err != nil {
			return fmt.Errorf("failed getting kubeconfig path: %w", err)
		}

		p := provider.NewInfraProvider(kubeconfig)

		err = builder.NewBuilder(p).BuildLock(builder.Model{
			Project:      name,
			Repo:         repo,
			URL:          url,
			ServivceName: service,
			Type:         types.Lock,
		},
		)

		return fmt.Errorf("failed building lock service: %w", err)
	}))

	create.AddCommand(newCommand("kvs", nil))
	create.AddCommand(newCommand("queue", nil))
	root.AddCommand(create)

	return root
}

func buildLockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "lock",
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

			err = builder.NewBuilder(p).BuildKVS(builder.Model{
				Project:      name,
				Repo:         repo,
				URL:          url,
				ServivceName: service,
				Type:         types.Lock,
			},
			)

			return fmt.Errorf("failed building lock service: %w", err)
		},
	}

	return cmd
}

func newBuildCommand(use string, f func(b *builder.Builder, m builder.Model) error) *cobra.Command {
	cmd := &cobra.Command{
		Use: use,
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
			b := builder.NewBuilder(p)
			err = f(b, builder.Model{
				Project:      name,
				Repo:         repo,
				URL:          url,
				ServivceName: service,
				Type:         types.Lock,
			},
			)

			return fmt.Errorf("failed building lock service: %w", err)
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

func runPut(cmd *cobra.Command, args []string) error {
	userID := args[0]
	game := args[1]

	id := guuid.New()

	s := infra.Server{Game: game, CreatedAt: time.Now(), GameID: id.String(), Owner: userID}

	if len(userID) == 0 || len(game) == 0 {
		return ErrorInvalidArgs
	}

	j, err := json.Marshal(s)
	if err != nil {
		return ErrorMarshalling
	}

	r := bytes.NewReader(j)

	resp, err := http.Post(url, contentType, r)
	if err != nil || resp.StatusCode != http.StatusCreated {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read server response")
	}

	fmt.Printf("%s", b)

	return nil
}
