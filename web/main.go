package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"flag"
)

func main() {
	ip := flag.String("ip", "", "IP address to use")
	port := flag.String("port", "8080", "Port to use")
	apiURL := flag.String("api", "http://dev-api.peppy.ml", "API address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "WEB! Version: %s Env: %s", os.Getenv("VERSION"), os.Getenv("APP_ENV"))
		fmt.Fprintf(w, "<a href='%s'>API Link</a>", *apiURL)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), nil))
}
