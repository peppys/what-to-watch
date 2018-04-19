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
	port := flag.String("port", "8081", "Port to use")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("API! Version: %s Env: %s", os.Getenv("VERSION"), os.Getenv("APP_ENV"))))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), nil))
}
