package services

import (
	"golang.org/x/net/context"
	"pb"
)

type EchoService struct {

}

func (this *EchoService) Echo(ctx context.Context, msg *pb.EchoMessage) (*pb.EchoMessage, error) {
	return msg, nil
}
