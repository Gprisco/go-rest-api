package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gprisco/nic-series-yt/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to encode json data", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusBadRequest)
		return
	}

	p.logger.Printf("Received: %#v", prod)
	data.AddProduct(prod)

	err = data.GetProducts().ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to Encode Response JSON", http.StatusInternalServerError)
		return
	}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Got ID: %v", id)

	// do something with the ID (update the given product)
}
