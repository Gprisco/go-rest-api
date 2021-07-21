package handlers

import (
	"log"
	"net/http"

	"github.com/gprisco/nic-series-yt/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		rw.Write([]byte("You just hit a PUT endpoint"))
		return
	}

	// catch every other request which is not GET
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to encode json data", http.StatusInternalServerError)
	}
}
