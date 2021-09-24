package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gprisco/nic-series-yt/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// GET Products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// POST Products
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// PUT Products
	if r.Method == http.MethodPut {
		regex := regexp.MustCompile(`/([0-9])+`)
		groups := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if groups == nil || len(groups[0]) == 0 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := groups[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Got ID: %v", id)

		return
	}

	// catch every other request which is not GET
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to encode json data", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusBadRequest)
		return
	}

	p.logger.Printf("Received: %#v", prod)
	data.AddProduct(prod)
}
