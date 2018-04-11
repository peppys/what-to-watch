package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
