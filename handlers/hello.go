package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the http.Handler interface
// https://pkg.go.dev/net/http#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World!")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s!", d)
}
