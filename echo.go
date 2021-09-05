package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func echoHandler(writer http.ResponseWriter, request *http.Request) {
	client_ip := request.Header.Get("x-forwarded-for")
	log.Println("Echoing request from:" + client_ip)

	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Errorf("Error reading body %w", err)
	}
	io.WriteString(writer, string(body))
}

func main() {
	log.Println("Starting echo server")
	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":8000", nil)
}
