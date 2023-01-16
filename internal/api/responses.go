package api

import (
	"net/http"
)

func CreateResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	CreateResponse(w, http.StatusNotFound, r.RequestURI+" Not Found")
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	CreateResponse(w, http.StatusMethodNotAllowed, r.Method+" is not allowed for endpoint "+r.RequestURI)
}

func InternalServerError(w http.ResponseWriter) {
	CreateResponse(w, http.StatusInternalServerError, "Internal Server Error")
}

func Success(w http.ResponseWriter, msg string) {
	CreateResponse(w, http.StatusOK, msg)
}
