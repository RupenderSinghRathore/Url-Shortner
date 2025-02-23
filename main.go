package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleHtml)

	fs := http.FileServer(http.Dir("templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates/", fs))

	mux.HandleFunc("/create", HandlePost)
	mux.HandleFunc("/{shortUrl}", HandleRedirect)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	log.Println("Listening on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
