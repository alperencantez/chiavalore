package server

import (
	"encoding/json"
	"net/http"

	"github.com/alperencantez/chiavalore/cache"
)

func CacheHandler(cache *cache.CV_Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			key := r.URL.Path[len("/cache/"):]
			value := cache.Get(key)
			jsonResponse(w, value)

		case "POST":
			var kv struct {
				Key   string      `json:"key"`
				Value interface{} `json:"value"`
			}

			if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cache.Set(kv.Key, kv.Value)
			w.WriteHeader(http.StatusCreated)

		case "DELETE":
			key := r.URL.Path[len("/cache/"):]
			cache.Delete(key)
			w.WriteHeader(http.StatusOK)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
