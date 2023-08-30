package main

import (
	"log"
	"net/http"
	"webpage-go/handlers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/ascii", handlers.Ascii)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}

}
