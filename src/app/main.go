package main

import (
	"net/http"
	"log"
	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr.DisableKeepAlives = true
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8002", proxy))
}
