package main

import  (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Prometheus endpoint
	http.Handle("/prometheus", promhttp.Handler())
	
	fmt.Println("Serving requests on port 3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

