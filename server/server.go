package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Service is running!")
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Calls: 50, Errors: 2")
	})
	http.ListenAndServe(":8080", nil)
}
