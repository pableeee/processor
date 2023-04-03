package steam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pableeee/processor/pkg/rest"
)

const (
	apiUrl = "https://api.steampowered.com/IGameServersService/GetServerList/v1/?key=%s&filter=%s&limit=100&appid=%s"
)

type QueryResponse struct {
	Response struct {
		Servers []struct {
			Addr       string `json:"addr"`
			Gameport   int    `json:"gameport"`
			Steamid    string `json:"steamid"`
			Name       string `json:"name"`
			Appid      int    `json:"appid"`
			Gamedir    string `json:"gamedir"`
			Version    string `json:"version"`
			Product    string `json:"product"`
			Region     int    `json:"region"`
			Players    int    `json:"players"`
			MaxPlayers int    `json:"max_players"`
			Bots       int    `json:"bots"`
			Map        string `json:"map"`
			Secure     bool   `json:"secure"`
			Dedicated  bool   `json:"dedicated"`
			Os         string `json:"os"`
		} `json:"servers"`
	} `json:"response"`
}

type Options struct {
	APIKey string
}

type ServersQuerier struct {
	client rest.Client[QueryResponse]
	Options
}

type QueryServerList struct {
	filters []string
	// Addr       *string `"tag:addr"`     // "100.67.132.7:27279",
	// Gameport   *string `"tag:gameport"` // 27279,
	// Steamid    *string `"tag:steamid"`  // "90171045896388612",
	// Name       *string `"tag:name"`     // "Valve CS:GO US West Server (srcds1005-eat1.117.265)",
	// AppID      string  `"tag:appid"`
	// Gamedir    *string `"tag:gamedir"`     // "csgo",
	// Version    *string `"tag:version"`     // "1.38.5.7",
	// Product    *string `"tag:product"`     // "csgo",
	// Region     *int    `"tag:region"`      // 1,
	// Players    *int    `"tag:players"`     // 0,
	// MaxPlayers *int    `"tag:max_players"` // 10,
	// Bots       *int    `"tag:bots"`        // 0,
	// Map        *string `"tag:map"`         // "de_dust2",
	// Secure     *bool   `"tag:secure"`      // true,
	// Dedicated  *bool   `"tag:dedicated"`   // true,
	// Os         *string `"tag:os"`          // "l",
	// Gametype   *string `"tag:gametype"`    // "valve_ds,empty,secure"
}

type QueryServerListOption func(*QueryServerList) *QueryServerList

var filterString = "\\%s\\%s"

func withFilter(qsl *QueryServerList, filter string) *QueryServerList {
	if qsl.filters == nil {
		qsl.filters = make([]string, 0)
	}
	qsl.filters = append(qsl.filters, filter)

	return qsl
}

func WithAppID(appID string) QueryServerListOption {
	return func(qsl *QueryServerList) *QueryServerList {
		return withFilter(qsl, fmt.Sprintf(filterString, "appid", appID))
	}
}

func WithGamedir(gamedir string) QueryServerListOption {
	return func(qsl *QueryServerList) *QueryServerList {
		return withFilter(qsl, fmt.Sprintf(filterString, "gamedir", gamedir))
	}
}

func WithMap(m string) QueryServerListOption {
	return func(qsl *QueryServerList) *QueryServerList {
		return withFilter(qsl, fmt.Sprintf(filterString, "map", m))
	}
}

func NewSteamServerQuerier(opt *Options) *ServersQuerier {
	return &ServersQuerier{
		client:  rest.NewRestClient[QueryResponse](),
		Options: *opt,
	}
}

func (s *ServersQuerier) getFilters(qsl *QueryServerList) string {
	return strings.Join(qsl.filters, "&")
}

func (s *ServersQuerier) QueryServerList(ctx context.Context, options ...QueryServerListOption) (*QueryResponse, error) {
	// Build the URL for the GetServerList method with the filter
	qsl := &QueryServerList{}
	for _, option := range options {
		option(qsl)
	}

	filter := s.getFilters(qsl)
	apiUrl := fmt.Sprintf(apiUrl, s.Options.APIKey, url.QueryEscape(filter))

	res, err := s.client.ExecuteWithContext(ctx, http.MethodGet, apiUrl, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed while querying API: %w", err)
	}

	return res, nil
}
