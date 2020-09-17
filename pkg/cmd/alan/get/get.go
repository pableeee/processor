package get

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	url = "http://127.0.0.1:8000/user/%s"
)

var ErrorInvalidArgs = fmt.Errorf("invalid number of arguments")

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "get",
		Short:   "Gets all games associated with a user",
		Long:    "Get all games associated with a user",
		Example: "alan get pableeee",
		RunE:    runGet,
	}

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	userID := args[0]

	if len(userID) == 0 {
		return ErrorInvalidArgs
	}

	resp, err := http.Get(fmt.Sprintf(url, userID))
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ErrorInvalidArgs
	}

	fmt.Printf("%s", b)

	return nil
}
