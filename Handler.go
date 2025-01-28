package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type postReq struct {
	Url string `json:"url"`
}
type resReq struct {
	ShortUrl string `json:"shortUrl"`
}

var urlMap = make(map[string]string)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}

	defer r.Body.Close()

	jsonReq := postReq{}
	if err := json.NewDecoder(r.Body).Decode(&jsonReq); err != nil {
		http.Error(w, "Couldn't decode json data", http.StatusBadRequest)
	}

	urlCode := CreatShortUrl(jsonReq.Url)

	urlMap[urlCode] = jsonReq.Url

	resStruct := resReq{ShortUrl: "http://localhost:8080/" + urlCode}

	if err := json.NewEncoder(w).Encode(resStruct); err != nil {
		http.Error(w, "Couldn't provide responce", http.StatusInternalServerError)
		log.Printf("Json encoding error: %v", err)
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("Couldn't get the ip address")
		return
	}
	log.Printf("Req completed for %v", ip)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	url, ok := urlMap[path]
	if !ok {
		http.Error(w, "Url not found in logs", http.StatusBadRequest)
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
