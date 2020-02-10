package main

import (
	"context"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/consul"

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
	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@127.0.0.1:5672"

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

	go func() {
		time.Sleep(time.Second * 2)
		r := proto.AddToQueueRequest{
			Name: "Peter",
		}
		resp, err := serviceBClient.AddToQueue(context.Background(), &r)
		if err != nil {
			panic(err)
		}
		log.Printf("Add message to the queue. MessageId: %s", resp.MessageId)
	}()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
