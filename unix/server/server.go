package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/steenzout/go-playground/unix/common"
)

// UnixDomainSocketHTTPServer implements a HTTP server over a unix domain socket connection.
type UnixDomainSocketHTTPServer struct {
	ln net.Listener
}

// ServeHTTP handle HTTP requests.
func (h *UnixDomainSocketHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var contents []byte

	contents, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("server: ERROR ", err)
		return
	}
	log.Println("server: contents=", bytes.Index(contents, []byte{0}))

	log.Println("server: r.URL.Scheme=", r.URL.Scheme)
	log.Println("server: r.URL.Host=", r.URL.Host)
	log.Println("server: r.URL.RawQuery=", r.URL.RawQuery)
	log.Println("server: r.URL.Path=", r.URL.Path)

	w.Write([]byte(fmt.Sprintf("r.URL.Path=", r.URL.Path, "\r\n")))
	w.Write([]byte(fmt.Sprintf("r.URL.Query=", r.URL.Query(), "\r\n")))
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
