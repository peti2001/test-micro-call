package subscriber

import (
	"context"
	"log"

	proto "github.com/peti2001/test-micro-call/serviceA/proto"
)

type Subscriber struct {
	ServiceBClient proto.ServiceBService
}

func (s Subscriber) Greeting(ctx context.Context, msg *proto.RabbitMQRequest) error {
	log.Printf("Consume message. Id: %s", msg.MessageId)

	req := &proto.AckMessageRequest{
		MessageId: msg.MessageId,
	}
	resp, err := s.ServiceBClient.AckMessage(ctx, req)
	//resp, err := s.ServiceBClient.AckMessage(context.Background(), req)
	if err != nil {
		panic(err)
	}
	log.Printf("Ack is sent to ServiceB. Status: %s\n", resp.Status)
	return nil
}
