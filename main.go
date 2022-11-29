package main

import (
	"log"
	"net/http"

	"www.github.com/N4choWasTaken/Go-API-JWT/login"
)

func main() {
	http.HandleFunc("/login", login.Login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
