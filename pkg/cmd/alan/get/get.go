package get

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	url = "http://127.0.0.1:8000/game/%s"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "get",
		Short:   "Gets all games asociated with a user",
		Long:    "Get all games asociated with a user",
		Example: "alan get pableeee",
		RunE:    runGet,
	}

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	userID := args[0]
	//cmd.Flags().StringVar(&userID, "user", "", "user to query")

	if len(userID) == 0 {
		return fmt.Errorf("invalid number of arguments")
	}

	resp, err := http.Get(fmt.Sprintf(url, userID))
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read server response")
	}

	fmt.Printf("%s", b)

	return nil
}
