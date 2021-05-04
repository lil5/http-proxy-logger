package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"sync/atomic"
)

// Request counter
var reqCounter int32
var wg sync.WaitGroup

type DebugTransport struct{}

func (DebugTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	wg.Add(1)
	defer wg.Done()
	atomic.AddInt32(&reqCounter, 1)

	requestDump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		return nil, err
	}
	log.Printf("---REQUEST %d---\n\n%s\n\n", reqCounter, string(requestDump))

	response, err := http.DefaultTransport.RoundTrip(r)
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		// copying the response body did not work
		return nil, err
	}

	log.Printf("---RESPONSE %d---\n\n%s\n\n", reqCounter, string(responseDump))
	return response, err
}

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

func getTarget() string {
	target := getEnv("TARGET", "http://example.com")
	return target
}

func main() {
	target, _ := url.Parse(getTarget())
	log.Printf("Forwarding %s -> %s\n", getListenAddress(), target)

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Transport = DebugTransport{}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// Update the headers to allow for SSL redirection
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = target.Host

		// Note that ServeHttp is non blocking and uses a go routine under the hood
		proxy.ServeHTTP(res, req)
	})

	/* d := proxy.Director
	proxy.Director = func(r *http.Request) {
		d(r) // call default director

		r.Host = target.Host // set Host header as expected by target
	}*/

	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
