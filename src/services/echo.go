package services

import (
	"golang.org/x/net/context"
	"pb"
	"fmt"
	"time"
)

type EchoService struct {

}

func echoMsg(msg string, count int) string {
	return fmt.Sprintf("Hello %s %d", msg, count)
}

func (this *EchoService) Echo(ctx context.Context, msg *pb.EchoMessage) (*pb.EchoMessage, error) {
	return &pb.EchoMessage{ Value: echoMsg(msg.Value, 0) }, nil
}

func (this *EchoService) StreamEcho(msg *pb.EchoMessage, stream pb.EchoService_StreamEchoServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&pb.EchoMessage{ Value: echoMsg(msg.Value, i)})
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
