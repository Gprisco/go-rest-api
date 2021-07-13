package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gprisco/nic-series-yt/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api => ", log.LstdFlags)
	hh := handlers.NewHello(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
