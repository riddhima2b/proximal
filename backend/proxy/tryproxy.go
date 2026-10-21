package proxy

import (
	"net/http/httputil"
	"net/url"
)

func TryProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(target)
			r.SetXForwarded()
		},
	}

	return proxy
}
