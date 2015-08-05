package client

import (
	"log"
	"net"
	"net/http"

	"github.com/steenzout/go-playground/unix/common"
)

// UnixDomainSocketHTTPClient implements an HTTP client over a unix domain socket connection.
type UnixDomainSocketHTTPClient struct {
	*http.Client
}

// NewUnixDomainSocketTransport creates a HTTP transport struct over a unix domain socket.
func NewUnixDomainSocketTransport(path string) *http.Transport {
	return &http.Transport{
		Dial: func(network string, address string) (net.Conn, error) {
			log.Println("client: network=", network, " address=", address)
			//-3 to strip out the :80 that gets added
			return net.Dial("unix", path)
		}}
}

// NewUnixDomainSocketHTTPClient creates a HTTP client to run over a unix domain socket connection.
func NewUnixDomainSocketHTTPClient(path string) *http.Client {
	var socket string
	if path == "" {
		socket = common.Domain_Socket
	} else {
		socket = path
	}

	return &http.Client{Transport: NewUnixDomainSocketTransport(socket)}
}
