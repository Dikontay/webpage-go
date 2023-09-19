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
			ErrorPage(w, 500)
			return
		}
		err = tm.Execute(w, nil)
		if err != nil {
			ErrorPage(w, 500)
			return
		}
	default:
		ErrorPage(w, http.StatusMethodNotAllowed)
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
			ErrorPage(w,  http.StatusInternalServerError)
			return
		}
		input := r.FormValue("userInput")
		font := r.FormValue("fonts")
		result, e := getformat.FinalOutput(input, font)
		fmt.Println(result)
		if !getformat.CheckLang(input) || strings.TrimSpace(input) == "" {
			ErrorPage(w,  400)
			return
		}
		if !e && result == "Bad Request" {
			ErrorPage(w,  500)
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
		ErrorPage(w,  405)
	}
}

