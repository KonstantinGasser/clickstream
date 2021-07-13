package main

import (
	"flag"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/cmd/httpserver"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/client"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", "localhost:8004", "address to run the server on")
	apptokenAddr := flag.String("apptoken-srv", "192.168.0.177:8006", "address to connect to app-token-service")
	appconfigAddr := flag.String("config-srv", "192.168.0.177:8005", "address to connect to app-config-service")
	flag.Parse()

	grpcAppToken, err := client.NewClientAppToken(*apptokenAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	grpcAppConfig, err := client.NewClientAppConfig(*appconfigAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	server := httpserver.NewDefault(*grpcAppToken, *grpcAppConfig)
	server.Apply(
		httpserver.WithAllowedOrgins("*"),
		httpserver.WithAllowedHeaders("x-datalab-token", "content-type"),
		httpserver.WithAllowedCreds,
	)

	server.Register("/api/v1/hello", server.Hello,
		server.WithCors,
		server.WithTraceIP,
		server.WithAuth,
	)
	server.Register("/api/v1/open", server.OpenSocket,
		server.WithTicketAuth,
	)

	logrus.Fatal(server.Start(*host))
}
