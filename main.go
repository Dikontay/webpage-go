package main

import (
	"fmt"
	"log"
	"net/http"
	"webpage-go/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/ascii", handlers.AsciiPage)
	fmt.Println("\nStarting with default port number : 8080")
	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
