package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "WEB! Version: %s Env: %s", os.Getenv("VERSION"), os.Getenv("APP_ENV"))
		fmt.Fprintf(w, "<a href='%s'>API Link</a>", os.Getenv("API_URL"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
