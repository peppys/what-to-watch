package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World! v5\n"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
