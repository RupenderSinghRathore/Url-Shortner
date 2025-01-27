package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", HandlePost)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
