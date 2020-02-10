package handler

import (
	"context"
	"log"
	"time"

	"github.com/micro/go-micro/client"
)
import proto "github.com/peti2001/test-micro-call/serviceB/proto"

type Handler struct {
	Client client.Client
}

func (h Handler) AddToQueue(ctx context.Context, req *proto.AddToQueueRequest, resp *proto.AddToQueueResponse) error {
	messageId := time.Now().Format(time.StampMicro)

	m := h.Client.NewMessage(
		"sayHello.topic",
		&proto.RabbitMQRequest{
			Name:      req.Name,
			MessageId: messageId,
		},
	)

	h.Client.Publish(ctx, m)

	log.Printf("Message has been added. Message: %s. Id: %s\n", req.Name, messageId)
	resp.MessageId = messageId

	return nil
}

func (h Handler) AckMessage(ctx context.Context, req *proto.AckMessageRequest, resp *proto.AckMessageResponse) error {
	log.Printf("Ack message. MessageId %s\n", req.MessageId)
	resp.Status = "ok"

	return nil
}
