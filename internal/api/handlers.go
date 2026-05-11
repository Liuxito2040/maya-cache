package api

import (
	"MayaCache/internal/backend"
	"MayaCache/internal/cache"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	Cache *cache.Cache
}

func NewHandler(c *cache.Cache) *Handler {
	return &Handler{
		Cache: c,
	}
}

func (h *Handler) GetCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	if value := h.Cache.Get(key); value != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"source": "cache",
			"value":  string(value),
		})

		return
	}

	value := backend.Fetch(key)
	h.Cache.Set(key, value, 2*time.Minute)
	json.NewEncoder(w).Encode(map[string]any{
		"source": "backend",
		"value":  string(value),
	})
}
