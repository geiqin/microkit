package main

import (
	"github.com/geiqin/microkit/kubernetes/config"
	k8s "github.com/geiqin/microkit/kubernetes/micro"
	"github.com/micro/go-micro/v2"
	"log"
)

type Say struct{}



func main() {
	service := k8s.NewService(
		micro.Name("greeter"),
	)

	service.Init()


	config.NewConfig()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
