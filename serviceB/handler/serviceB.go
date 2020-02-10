package handler

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
)
import proto "github.com/peti2001/test-micro-call/serviceB/proto"

type Handler struct {
	Client client.Client
}

func (h Handler) AckMessage(ctx context.Context, req *proto.AckMessageRequest, resp *proto.AckMessageResponse) error {
	log.Printf("Ack recieved of messageId %s\n", req.MessageId)
	resp.Status = "ok"

	return nil
}
