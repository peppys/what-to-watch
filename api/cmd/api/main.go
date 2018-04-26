package main

import (
	"net/http"
	"log"
	"fmt"
	"flag"
	"time"
	"github.com/PeppyS/personal-site-api/api/router"
)

func main() {
	ip := flag.String("ip", "", "IP address to use")
	port := flag.String("port", "8081", "Port to use")
	flag.Parse()

	srv := &http.Server{
		Handler: router.New(),
		Addr: fmt.Sprintf("%s:%s", *ip, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
