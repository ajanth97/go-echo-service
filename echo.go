package main

import (
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	client_ip := r.Header.Get("x-forwarded-for")
	log.Println("Echoing request from" + client_ip)
	r.Write(w)
}

func main() {
	log.Println("Starting echo server")
	http.HandleFunc("/", echoHandler)
}
