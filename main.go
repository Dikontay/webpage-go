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

//func renderTemplate(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles(tmpl)
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/ascii", handlers.Ascii)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}

}
