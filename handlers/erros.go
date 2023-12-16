package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type Error struct{
	Code int
	Text string
}

func ErrorPage(w http.ResponseWriter,  status int) {
	w.WriteHeader(status)
	ts, err := template.ParseFiles("./templates/errors.html")
	if err != nil{
		fmt.Println(err)
		http.Error(w,  http.StatusText(status), status)
		return
	}
	e := &Error{
		Code: status,
		Text: http.StatusText(status),
	}
	err = ts.Execute(w, e)
	if err!= nil {
		fmt.Println(err)
		http.Error(w,  http.StatusText(status), status)
	}
}
