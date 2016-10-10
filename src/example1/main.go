package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"pb"
	"services"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	gRPCPort        = flag.Int("grpc_port", 9000, "port of gRPC service")
	gRPCGatewayPort = flag.Int("grpc_gateway", 9001, "port of gRPC gateway")
)

func runGRPC() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCPort))
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &services.EchoService{})

	s.Serve(l)
}

func runGateway() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEchoServiceHandlerFromEndpoint(context.TODO(), mux, fmt.Sprintf(":%d", *gRPCPort), opts)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *gRPCGatewayPort), mux)
	return
}

func main() {
	flag.Parse()

	go runGRPC()

	runGateway()
}
