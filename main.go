package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, HTMX!"))
}

func main() {
	// Handle the hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
