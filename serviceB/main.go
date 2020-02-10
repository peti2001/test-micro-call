package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/consul"

	"github.com/peti2001/test-micro-call/serviceB/handler"
	proto "github.com/peti2001/test-micro-call/serviceB/proto"
)

func main() {

	amqpBroker := rabbitmq.NewBroker()
	serviceName := "serviceB"

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Broker(amqpBroker),
		micro.RegisterTTL(15*time.Second),
		micro.RegisterInterval(5*time.Second),
	)

	service.Init()
	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@127.0.0.1:5672"

	// HANDLERS
	h := &handler.Handler{
		Client: service.Client(),
	}
	proto.RegisterServiceBHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
