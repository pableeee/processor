package put

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	guuid "github.com/google/uuid"
	"github.com/pableeee/processor/pkg/cmd/processor/infra"
	"github.com/spf13/cobra"
)

const (
	url         = "http://127.0.0.1:8000/game/"
	contentType = "application/json"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(2),
		Use:     "put",
		Short:   "Associates a game to a user",
		Long:    "Associates a game to a user",
		Example: "alan put USER GAME",
		RunE:    runPut,
	}

	return cmd
}

func runPut(cmd *cobra.Command, args []string) error {
	userID := args[0]
	game := args[1]

	id := guuid.New()

	s := infra.Server{Game: game, CreatedAt: time.Now(), GameID: id.String(), Owner: userID}

	if len(userID) == 0 || len(game) == 0 {
		return fmt.Errorf("invalid arguments")
	}

	j, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("unable to marshal info to backend")
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
