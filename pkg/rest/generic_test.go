package rest

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedRestClient struct {
	mock.Mock
}

func (g *mockedRestClient) Do(req *http.Request) (*http.Response, error) {
	args := g.Called(req)

	return args.Get(0).(*http.Response), args.Error(1)
}

var stringResponse = `
{
	"name": "release: v0.371",
	"tag_name": "v0.371",
	"description": "\n# Release: v0.371\ncommit: cd4f174512b4e2a0f321d1077580af1dfb04de57\nauthor: Pablo Duco <pablo.duco@wildlifestudios.com>\n## public\n- marta.attribution.production.yaml\n- marta.attribution.staging.yaml\n- marta.autoloader.production.yaml\n- marta.autoloader.staging.yaml\n        ",
	"commit": {
		"id": "cd4f174512b4e2a0f321d1077580af1dfb04de57",
		"short_id": "cd4f1745",
		"title": "Merge branch 'Update-README' into 'main'",
		"message": "Merge branch 'Update-README' into 'main'\n\nUpdate Readme\n\nSee merge request attribution/config-repository-template!12"
	}
}
`

type Release struct {
	Name        string `json:"name"`
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
	Commit      struct {
		ID      string `json:"id"`
		ShortID string `json:"short_id"`
		Title   string `json:"title"`
		Message string `json:"message"`
	} `json:"commit"`
}

func TestParseHttpResponse(t *testing.T) {
	// Release represents a gitlab api release object.
	mc := &mockedRestClient{}
	client := NewRestClient[Release]()
	client.client = mc

	mc.On("Do", mock.Anything).
		Return(&http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(stringResponse)),
		}, nil)

	s, err := client.Execute("GET", "url", nil, nil)
	assert.Nil(t, err)

	fmt.Printf("%+v", s)

}
