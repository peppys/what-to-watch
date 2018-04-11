package main

import (
	"net/http"
	"log"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})

	// Redirect for https
	go http.ListenAndServe(":80", http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		targetURL := url.URL{ Scheme: "https", Host: strings.Split(req.Host, ":")[0] + ":443", Path: req.URL.Path, RawQuery: req.URL.RawQuery, }
		http.Redirect(w, req, targetURL.String(), http.StatusTemporaryRedirect)
	}))

	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
