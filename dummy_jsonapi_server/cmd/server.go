package main

import (
	"fmt"
	"net/http"

	"github.com/dailypay/jsonapitest/internal"
)

// start server and serve the objects
func main() {
	handler := &internal.RequestHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.ServeHTTP)

	fmt.Printf("Starting server at port 9999\n")
	http.ListenAndServe(":9999", mux)
}
