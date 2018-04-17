package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("API! Version: %s Env: %s", os.Getenv("VERSION"), os.Getenv("APP_ENV"))))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
