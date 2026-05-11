package main

import (
	"MayaCache/internal/api"
	"MayaCache/internal/cache"
	"log"
	"net/http"
)

func main() {
	cache := cache.New()
	handler := api.NewHandler(cache)
	http.HandleFunc("/cache", handler.GetCache)
	log.Println("cache code running on: ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
