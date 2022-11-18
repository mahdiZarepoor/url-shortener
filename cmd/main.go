package main

import (
	"net/http"

	"github.com/mahdiZarepoor/url-shortener/internal/handlers"
)


func main() {
	myApp := handlers.NewApp()
	panic(http.ListenAndServe("localhost:8080", myApp.Router))
}