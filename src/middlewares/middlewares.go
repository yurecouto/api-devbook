package middlewares

import (
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Printf("\n %s %s %s", request.Method, request.RequestURI, request.Host)
		next(response, request)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		next(response, request)
	}
}
