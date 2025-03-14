package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Service is running!")
	})
	log.Println("Server started on port 8080...")
	err := http.ListenAndServe(":8080", nil) 
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
