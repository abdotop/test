package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
