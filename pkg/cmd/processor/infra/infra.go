package infra

import "time"

// Server represents a game server
type Server struct {
	Owner     string    `json:"owner"`
	Game      string    `json:"game"`
	GameID    string    `json:"id"`
	CreatedAt time.Time `json:"created-at"`
}

// Backend represents the backend server storing
type Backend interface {
	Get(userID string) ([]Server, error)
	Put(userID string, s Server) error
}
