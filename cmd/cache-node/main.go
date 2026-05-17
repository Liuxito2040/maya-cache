package main

import (
	"MayaCache/internal/api"
	"MayaCache/internal/cache"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	c := cache.New()
	handler := api.NewHandler(c)
	http.HandleFunc("/internal/cache", handler.GetCache)

	log.Println("cache node running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
