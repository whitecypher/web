package webserve

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

type server struct {
	started bool
	done    chan bool
	err     error
	addr    string
	server  *http.Server
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

// Start the webserver
func (ws *server) Start() *server {
	go ws.start_concurrent()
	return ws
}

// Get the error given by webserver on exit
func (ws *server) Err() error {
	return ws.err
}

// Internal webserver starter to be used as goroutine
func (ws *server) start_concurrent() {
	ws.err = gracehttp.Serve(ws.server)
	ws.done <- true
}

// Wait the server to finish
func (ws *server) Wait() bool {
	return <-ws.done
}
