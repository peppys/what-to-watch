package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PeppyS/what-to-watch/api/controller"
	"github.com/PeppyS/what-to-watch/api/server"
	"github.com/PeppyS/what-to-watch/api/service"
)

func main() {
	ip := flag.String("ip", "", "IP address to use")
	port := flag.String("port", "8080", "Port to use")

	flag.Parse()

	httpAddress := fmt.Sprintf("%s:%s", *ip, *port)

	healthAPI := controller.NewHealth()
	movieAPI := controller.NewMovie(
		service.NewMovie(),
	)

	go server.ListenAndServe(":50051", movieAPI, healthAPI)

	log.Fatal(server.ListenAndServeHTTPGateway(":50051", httpAddress))
}
