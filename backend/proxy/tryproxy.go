package proxy

import (
	"net/http/httputil"
	"net/url"
)

func TryProxy() *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:9000",
	})
	return proxy
}
