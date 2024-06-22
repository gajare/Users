package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started :%s ,\t %s ", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s ,\t%s in \t %v", r.Method, r.RequestURI, time.Since(start))
	})
}
