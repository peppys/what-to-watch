package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PeppyS/personal-site-api/controller"
	"github.com/PeppyS/personal-site-api/repository"
	"github.com/PeppyS/personal-site-api/server"
	"github.com/PeppyS/personal-site-api/service"
)

func main() {
	ip := flag.String("ip", "", "IP address to use")
	port := flag.String("port", "8081", "Port to use")

	flag.Parse()

	httpAddress := fmt.Sprintf("%s:%s", *ip, *port)

	healthAPI := controller.NewHealth()
	resumeAPI := controller.NewResume(
		service.NewResume(
			repository.NewResume(),
		),
	)

	go server.ListenAndServe(":50051", resumeAPI, healthAPI)

	log.Fatal(server.ListenAndServeHTTPGateway(":50051", httpAddress))
}
