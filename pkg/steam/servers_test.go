package steam

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryServers(t *testing.T) {
	// Release represents a gitlab api release object.
	q := NewSteamServerQuerier(&Options{APIKey: "F053B6FA7FF33FFF64972A37FAD5AF56"})

	res, err := q.QueryServerList(context.TODO(), "10", "\\map\\de_dust2")
	assert.Nil(t, err)

	// Print out the IP address and port of each server in the list
	for _, server := range res.Response.Servers {
		fmt.Printf("%+v", server)
	}
}
