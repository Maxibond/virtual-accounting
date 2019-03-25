package main

import (
	"log"
	"net/http"
	"time"
)

var app = NewApplication()

func main() {

	http.HandleFunc("/accounts", HandleCreateAccount)
	http.HandleFunc("/balance", HandleGetAccountBalance)
	http.HandleFunc("/send", HandleCreateMove)

	s := &http.Server{
		Addr:           ":8000",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
