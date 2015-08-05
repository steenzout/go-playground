package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/steenzout/go-playground/unix/common"
	"log"
)

// UnixDomainSocketHTTPServer implements a HTTP server over a unix domain socket connection.
type UnixDomainSocketHTTPServer struct {
	ln net.Listener
}

// ServeHTTP handle HTTP requests.
func (h *UnixDomainSocketHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("server: r.URL.Scheme=", r.URL.Scheme)
	log.Println("server: r.URL.Host=", r.URL.Host)
	log.Println("server: r.URL.RawQuery=", r.URL.RawQuery)
	log.Println("server: r.URL.Path=", r.URL.Path)

	fmt.Fprintf(w, "r.URL.Path=", r.URL.Path, "\r\n")
	fmt.Fprintf(w, "r.URL.Query=", r.URL.Query(), "\r\n")

}

// Close closes the unix domain socket connection.
func (h *UnixDomainSocketHTTPServer) Close() {
	h.ln.Close()
}

// Serve handle HTTP requests.
func (h *UnixDomainSocketHTTPServer) Serve() {
	var err error

	log.Println("server: starting")
	if err = http.Serve(h.ln, h); err != nil {
		h.ln.Close()
	}
}

// NewUnixDomainSocketServer creates a new HTTP server over unix domain socket.
func NewUnixDomainSocketServer(path string) *UnixDomainSocketHTTPServer {
	var ln net.Listener
	var err error
	var socket string

	if path == "" {
		socket = common.Domain_Socket
	} else {
		socket = path
	}

	if ln, err = net.Listen("unix", socket); err != nil {
		panic(err)
	}
	return &UnixDomainSocketHTTPServer{ln}
}
