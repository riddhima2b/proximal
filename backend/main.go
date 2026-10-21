package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"proximal/proxy"
)

func main() {

	port := flag.String("port", "", "Port to listen on")
	origin := flag.String("origin", "", "Origin server to proxy to")

	flag.Parse()

	if *origin == "" {
		log.Fatal("Origin server must be specified")
	}
	target, err := url.Parse(*origin)
	if err != nil {
		log.Fatalf("Failed to parse origin URL: %v", err)
	}
	p := proxy.TryProxy(target)

	addr := ":" + *port

	http.Handle("/", p)

	log.Printf("Proxy listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
