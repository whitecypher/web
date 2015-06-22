package webserve

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

type server struct {
	started bool
	done    chan error
	addr    string
	server  *http.Server
}

// Create a new webserver
func New(addr string, handler http.Handler) *server {
	ws := &webserver{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		done: make(chan error),
	}

	return ws
}

// Start the webserver
func (ws *server) Start() *server {
	go ws.start_concurrent()
	return ws
}

// Internal webserver starter to be used as goroutine
func (ws *server) start_concurrent() {
	ws.done <- gracehttp.Serve(ws.server)
}

// Wait the server to finish
func (ws *server) Wait() error {
	return <-ws.done
}
