package middleware

import (
	"log"
	"net/http"
	"time"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Authorizing...")
		h.ServeHTTP(writer, request)
	})
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(writer, request)
		log.Printf("%s - %s (%v)\n", request.Method, request.URL.Path, time.Since(startTime))
	})
}
