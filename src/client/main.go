package main

import (
	"flag"
	"fmt"

	"pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("p", 9000, "port of gRPC service")
	echo = flag.String("e", "", "message for echo service")
	v0 = flag.Int("v0", 0, "1st number for calculate service")
	v1 = flag.Int("v1", 0, "2nd number for calculate service")
	op = flag.String("op", "", "operator for calculate service")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if len(*echo) != 0 {
		client := pb.NewEchoServiceClient(conn)
		r, err := client.Echo(context.TODO(), &pb.EchoMessage{
			Value: *echo,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", r)
	} else {
		var opr pb.CalculateRequestOp
		switch *op {
		case "ADD":
			fallthrough
		case "add":
			opr = pb.CalculateRequest_ADD
		case "sub":
			fallthrough
		case "SUB":
			opr = pb.CalculateRequest_SUB
		case "MUL":
			fallthrough
		case "mul":
			opr = pb.CalculateRequest_MUL
		case "DIV":
			fallthrough
		case "div":
			opr = pb.CalculateRequest_DIV
		default:
			panic(fmt.Sprintf("Invalid op: %s", *op))
		}

		client := pb.NewCalculateServiceClient(conn)
		r, err := client.Calculate(context.TODO(), &pb.CalculateRequest{
			Value0: int32(*v0),
			Value1: int32(*v1),
			Operator: opr,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", r)
	}
}
