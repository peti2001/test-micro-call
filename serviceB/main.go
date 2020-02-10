package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/consul"

	"github.com/peti2001/test-micro-call/serviceB/handler"
	proto "github.com/peti2001/test-micro-call/serviceB/proto"
)

func main() {
	serviceName := "serviceB"

	service := micro.NewService(
		micro.Name(serviceName),
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

	go func() {
		time.Sleep(time.Second * 2)
		messageId := time.Now().Format(fmt.Sprintf("%d", time.Now().Nanosecond()))

		m := h.Client.NewMessage(
			"sayHello.topic",
			&proto.RabbitMQRequest{
				Name:      "Peter",
				MessageId: messageId,
			},
		)

		h.Client.Publish(context.Background(), m)

		log.Printf("Ask ServiceA to do a long process by publishing a message. Waiting for ack so I can continue after that. MessageId: %s\n", messageId)
	}()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
