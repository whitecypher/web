package webserve

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

type server struct {
	addr   string
	server *http.Server
	err    error
	done   chan bool
}

// Create a new webserver
func New(addr string, handler http.Handler) *server {
	ws := &server{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		done: make(chan bool),
	}

	return ws
}

// Error returned by webserver exit
func (ws *server) Err() error {
	return ws.err
}

// Start the webserver
func (ws *server) Start() *server {
	go ws.start_concurrent()
	return ws
}

// Internal webserver starter to be used as goroutine
func (ws *server) start_concurrent() {
	ws.err = gracehttp.Serve(ws.server)
	close(ws.done)
}

// Channel to wait on
func (ws *server) Done() <-chan bool {
	return ws.done
}
