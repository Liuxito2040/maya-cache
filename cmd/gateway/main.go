package main

import (
	"MayaCache/internal/ring"
	"log"
	"net/http"
)

func main() {
	r := ring.New()

	r.AddNode("http://localhost:8081")
	r.AddNode("http://localhost:8082")
	r.AddNode("http://localhost:8083")

	handler := NewGatewayHandler(r)

	http.HandleFunc("/cache", handler.Get)
	log.Println("gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
