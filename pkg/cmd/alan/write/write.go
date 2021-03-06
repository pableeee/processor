package write

import (
	"fmt"

	"github.com/pableeee/processor/pkg/queue"
	"github.com/spf13/cobra"
)

var (
	ErrorInvalidArgs = fmt.Errorf("invalid arguments")
	ErrorMarshalling = fmt.Errorf("unable to marshal info to backend")
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(2),
		Use:     "write",
		Short:   "write a message into a nats queue",
		Long:    "write a message into a nats queue",
		Example: "alan write topic message",
		RunE:    runWrite,
	}

	return cmd
}

func runWrite(cmd *cobra.Command, args []string) error {
	topic := args[0]
	msg := args[1]

	p, err := queue.NewNatsPublisher("127.0.0.1", 5555)

	if err != nil {
		return err
	}
	defer p.Close()

	err = p.Publish(topic, []byte(msg))
	if err != nil {
		return err
	}

	fmt.Println("message published")

	return nil
}
