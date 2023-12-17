package handler

import "net/http"

func HandleNotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(404)
	w.Write([]byte("route does not exist"))
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(405)
	w.Write([]byte("method is not valid"))

}
