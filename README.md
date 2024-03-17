# Chiavalore - A Simple Key-Value Memory Cache Library in Go

Chiavalore is a lightweight Go library that provides a simple key-value memory cache implementation. It offers basic functionalities for storing, retrieving, and deleting key-value pairs in memory. The library also includes an HTTP server component for interacting with the cache over HTTP.

## Features

- Easy-to-use API for setting, getting, and deleting key-value pairs.
- Support for concurrent access with built-in synchronization using mutex.
- Simple HTTP server for interacting with the cache over HTTP.

## Installation

To use Chiavalore in your Go project, you can install it via `go get`:

```bash
go get github.com/alperencantez/chiavalore
```

## Usage

```go
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

```

## Contribution

Contributions to Chiavalore are welcome! If you find any bugs, have feature requests, or want to contribute improvements, feel free to open an issue or submit a pull request on GitHub.
