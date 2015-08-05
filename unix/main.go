package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/steenzout/go-playground/unix/client"
	"github.com/steenzout/go-playground/unix/server"
)

func main() {
	var res *http.Response
	var contents []byte
	var err error

	server := server.NewUnixDomainSocketServer("")
	go server.Serve()
	defer server.Close()

	client := client.NewUnixDomainSocketHTTPClient("")
	log.Println("main: GET http://www.domain.com/path?param=value")
	if res, err = client.Get("http://www.domain.com/path?param=value"); err != nil {
		log.Println("client: ERROR ", err)
		os.Exit(1)
	}

	defer res.Body.Close()
	contents, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("client: ERROR ", err)
		os.Exit(1)
	}
	log.Println("client: contents=", bytes.Index(contents, []byte{0}))

	log.Println()

	log.Println("HEAD http://www.domain.com/head")
	if res, err = client.Head("http://www.domain.com/head"); err != nil {
		log.Println("client: ERROR ", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	contents, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("client: ERROR ", err)
		os.Exit(1)
	}
	log.Println("client: contents=", bytes.Index(contents, []byte{0}))
}
