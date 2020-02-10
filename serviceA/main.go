package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro/v2"
	proto "github.com/peti2001/test-micro-call/serviceA/proto"
	"github.com/peti2001/test-micro-call/serviceA/subscriber"
)

func main() {
	serviceName := "serviceA"

	service := micro.NewService(
		micro.Name(serviceName),
		micro.RegisterTTL(15*time.Second),
		micro.RegisterInterval(5*time.Second),
	)

	service.Init()

	//CLIENT
	serviceBClient := proto.NewServiceBService("serviceB", service.Client())

	//SUBSCRIBERS
	sub := service.Server().NewSubscriber(
		"sayHello.topic",
		subscriber.Subscriber{
			ServiceBClient: serviceBClient,
		},
	)
	service.Server().Subscribe(sub)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
