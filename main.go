package main

import (
	"fmt"
	"net/http"

	"github.com/alperencantez/chiavalore/cache"
	"github.com/alperencantez/chiavalore/server"
)

func main() {
	const PORT = 8080
	var cv_cache *cache.CV_Cache = cache.InitCV_Cache() // initialize Chiavalore

	// Route definition for handler server
	http.HandleFunc("/cache/", server.CacheHandler(cv_cache))

	fmt.Printf("Chiavalore mem-cache server is up and running on %d", PORT)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
