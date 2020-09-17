package del

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	url         = "http://127.0.0.1:8000/game/%s"
	contentType = "application/json"
)

var ErrorInvalidArgs = fmt.Errorf("invalid arguments")

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "del",
		Short:   "Deletes a game from a user",
		Long:    "Associates a game to a user",
		Example: "alan del <GAME_ID>",
		RunE:    runDel,
	}

	return cmd
}

func runDel(cmd *cobra.Command, args []string) error {
	gameID := args[0]

	if len(gameID) == 0 {
		return ErrorInvalidArgs
	}

	// Create client
	client := &http.Client{}
	uri := fmt.Sprintf(url, gameID)
	// Create request
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		log.Fatal(err.Error())

		return err
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal(err.Error())

		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Game %s was successfully terminated", gameID)

	return nil
}
