package main

import (
	"net/http"
	"net/http/httputil"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		httputil.NewSingleHostReverseProxy(r.URL).ServeHTTP(w, r)
	})
	http.ListenAndServe(":8002", nil)
}

