package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	port := flag.String("port", "8080", "Port to listen on")
	origin := flag.String("origin", "", "Origin server to proxy to")

	flag.Parse()

	if *origin == "" {
		log.Fatal("Origin server must be specified")
	}
	target, err := url.Parse(*origin)
	if err != nil {
		log.Fatalf("Failed to parse origin URL: %v", err)
	}
	p := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(target)
			r.SetXForwarded()
		},
	}

	addr := ":" + *port

	http.Handle("/", p)

	log.Printf("Proxy listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
