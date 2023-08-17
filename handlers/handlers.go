package handlers

import (
	"html/template"
	"net/http"
	"webpage-go/getFormat"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		tm, err := template.ParseFiles("/Users/dilnaz/Desktop/go learning/webpage-go/templates/index.html")
		if err != nil {
			errorPage(w, 500)
			return
		}
		err = tm.Execute(w, nil)
		if err != nil {
			errorPage(w, 500)
			return
		}
	default:
		errorPage(w, http.StatusMethodNotAllowed)
	}
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "POST" {
		tm, err := template.ParseFiles("/Users/dilnaz/Desktop/go learning/webpage-go/templates/index.html")
		if err != nil {
			errorPage(w, 500)
			return
		}
		userInput := r.FormValue("userInput")
		format := r.FormValue("fonts")
		res, okay := getformat.FinalOutput(userInput, format)
		if !okay {
			errorPage(w, http.StatusInternalServerError)
			return
		}
		Result := struct {
			Word string
			Res  string
		}{
			Word: userInput,
			Res:  res,
		}
		err = tm.Execute(w, Result)
		if err != nil {
			errorPage(w, 500)
			return
		}
	}
}
