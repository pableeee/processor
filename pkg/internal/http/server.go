package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Builder struct {
	addr    string
	port    int
	writeTO time.Duration
	readTO  time.Duration
	setup   func(r *mux.Router)
}

func DefaultBuilder() *Builder {
	r := Builder{
		addr: "127.0.0.1",
		port: 8000,
		// Good practice: enforce timeouts for servers you create!
		writeTO: 15 * time.Second,
		readTO:  15 * time.Second,
	}

	return &r
}

func (b *Builder) WithAddress(addr string) *Builder {
	b.addr = addr

	return b
}

func (b *Builder) WithPort(p int) *Builder {
	b.port = p

	return b
}

func (b *Builder) WithWriteTimeout(wto time.Duration) *Builder {
	b.writeTO = wto

	return b
}

func (b *Builder) WithReadTimeout(rto time.Duration) *Builder {
	b.readTO = rto

	return b
}

// f function should do all the HandleFuncs.
func (b *Builder) WithHandlerSetUp(f func(r *mux.Router)) *Builder {
	b.setup = f

	return b
}

func (b *Builder) Build() *http.Server {
	r := mux.NewRouter()

	if b.setup != nil {
		b.setup(r)
	}

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", b.addr, b.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: b.writeTO,
		ReadTimeout:  b.readTO,
	}

	return srv
}
