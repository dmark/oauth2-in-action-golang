// a minimalist "echo" webserver
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc registers a handler function for the given pattern
	http.HandleFunc("/", handler) // each request calls handler
	// Start the server on port 9000
	fmt.Println("Server is starting on port 9000...")
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// handler echos the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
