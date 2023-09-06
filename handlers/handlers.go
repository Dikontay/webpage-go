package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"webpage-go/getFormat"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		tm, err := template.ParseFiles("./templates/index.html")
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

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST":
		ts, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			ErrorPage(w, r, http.StatusInternalServerError)
			return
		}
		input := r.FormValue("userInput")
		font := r.FormValue("fonts")
		result, e := getformat.FinalOutput(input, font)
		fmt.Println(result)
		if !getformat.CheckLang(input) || strings.TrimSpace(input) == "" {
			ErrorPage(w, r, 400)
			return
		}
		if !e && result == "Bad Request" {
			ErrorPage(w, r, 500)
			return
		}
		err = ts.Execute(w, struct {
			Result string
			Word   string
		}{
			Result: result,
			Word:   input,
		})
	default:
		ErrorPage(w, r, 405)
	}
}

func ErrorPage(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	switch status {
	case 404:
		w.Write([]byte("Page Not Found"))
	case 405:
		w.Write([]byte("Method Not Allowed"))
	case 400:
		w.Write([]byte("Bad Request"))
	case 500:
		w.Write([]byte("Internal Server Error"))
	default:
		return
	}
}
