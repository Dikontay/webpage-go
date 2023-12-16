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
	mux.HandleFunc("/download", handlers.Download)
	fmt.Println("Starting the server: http://localhost:8000/")
	err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}
