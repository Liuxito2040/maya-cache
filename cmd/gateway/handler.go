package main

import (
	"MayaCache/internal/ring"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type GatewayHandler struct {
	ring *ring.Ring
}

func NewGatewayHandler(r *ring.Ring) *GatewayHandler {
	return &GatewayHandler{
		ring: r,
	}
}

func (h *GatewayHandler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	node := h.ring.GetNode(key)
	log.Printf("key=%s routed_to=%s", key, node)

	resp, err := http.Get(node + "/internal/cache?key=" + key)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"node":     node,
		"response": json.RawMessage(body),
	})
}
