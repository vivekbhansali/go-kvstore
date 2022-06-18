package main

import (
	"log"
	"net/http"

	"github.com/vivekbhansali/go-kvstore/handler"
)

func main() {
	handler := handler.New()
	server := &http.Server{
		Addr: ":8080",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}