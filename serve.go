package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("public"))
	http.Handle("/", server)

	log.Println("Listening...")
	http.ListenAndServe(":8181", nil)
}
