package handlers

import "net/http"

func errorPage(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	switch statusCode {
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
