package subscriber

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/peti2001/test-micro-call/serviceA/proto"
)

type Subscriber struct {
	ServiceBClient proto.ServiceBService
}

func (s Subscriber) Greeting(ctx context.Context, msg *proto.RabbitMQRequest) error {
	log.Printf("Consume message. MessageId: %s ", msg.MessageId)
	log.Printf("Hello, %s\n", msg.Name)
	time.Sleep(time.Second * 2)
	//DEBUG will output this:
	//&context.valueCtx{
	//	Context:(*context.emptyCtx)(0xc0000420b0), key:metadata.metaKey{}, val:metadata.Metadata{
	//		"Accept":"application/protobuf",
	//	 	"Content-Length":"7", "Content-Type":"application/protobuf",
	//	 	"Local":"[::]:61941",
	//	 	"Micro-Endpoint":"ServiceB.AddToQueue",
	//	 	"Micro-From-Service":"serviceA",
	//	 	"Micro-Id":"ff54ecf8-7826-4d1c-aeaa-11102da736e3",
	//	 	"Micro-Method":"ServiceB.AddToQueue",
	//	 	"Micro-Service":"serviceB",
	//	 	"Micro-Topic":"sayHello.topic",
	//	 	"Remote":"192.168.100.2:62195",
	//	 	"Timeout":"5000000000",
	//	 	"User-Agent":"Go-http-client/1.1"
	//	}
	//}
	//
	// The problem here is the "Micro-Topic":"sayHello.topic", which was added after
	// the message was published and was carried by the context.
	fmt.Printf("DEBUG ctx %#v \n", ctx)

	req := &proto.AckMessageRequest{
		MessageId: msg.MessageId,
	}
	// This call won't work
	resp, err := s.ServiceBClient.AckMessage(ctx, req)
	log.Printf("Ack is sent to ServiceB. Status: %s\n", resp.Status)

	// This call will work
	resp, err = s.ServiceBClient.AckMessage(context.Background(), req)
	log.Printf("Ack is sent to ServiceB with new context. Status: %s\n", resp.Status)
	if err != nil {
		panic(err)
	}
	return nil
}
