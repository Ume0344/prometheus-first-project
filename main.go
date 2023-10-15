package main

import  (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"html/template"
)

// Create a var of prometheus Counter
var httpRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
)

// Increment the counter whenever :3000 is accessed
func recordHttpRequests(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./static/index.html")
	t.Execute(w, "")
	httpRequests.Inc()
}

// Register the httRequest counter
func init() {
	prometheus.Register(httpRequests)
}

func main() {
	http.HandleFunc("/", recordHttpRequests)
	// Prometheus endpoint
	http.Handle("/prometheus", promhttp.Handler())
	
	fmt.Println("Serving requests on port 3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
