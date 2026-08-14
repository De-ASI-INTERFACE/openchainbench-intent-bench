package main

// Minimal harness skeleton for intent-solver-latency benchmark.
// TODO: implement quote/fill loops, Prometheus metrics, and provider clients.

import (
	"fmt"
	"net/http"
	_ "github.com/prometheus/client_golang/prometheus"
)

func main() {
	http.Handle("/metrics", http.NotFoundHandler()) // TODO: wire real metrics
	fmt.Println("Listening on :8080/metrics")
	http.ListenAndServe(":8080", nil)
}
