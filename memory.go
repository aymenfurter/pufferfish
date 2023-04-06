package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	// 1 GB in bytes
	oneGigabyte = 1024 * 1024 * 1024
)

func main() {
	// Allocate approximately 1 GB of memory
	mem := make([]byte, oneGigabyte)
	// Touch the memory to ensure it's truly allocated
	for i := range mem {
		mem[i] = byte(i % 256)
	}

	// Set up HTTP server for monitoring
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Memory test app, allocated 1 GB of RAM.")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting memory test app, allocated 1 GB of RAM, listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
