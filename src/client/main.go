package main

import (
	"flag"
	"fmt"

	"pb"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
)


var (
	port        = flag.Int("grpc_port", 9000, "port of gRPC service")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	r, err := client.Echo(context.TODO(), &pb.EchoMessage{
		Value: "foo",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", r)
}