package main

import (
	"fmt"
	"net"
	"net/http"

	"pb"
	"services"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func runGRPC() {
	l, err := net.Listen("tcp", ":9090")
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
	err := pb.RegisterEchoServiceHandlerFromEndpoint(context.TODO(), mux, "localhost:9090", opts)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	http.ListenAndServe(":8080", mux)
	return
}

func main() {

	go runGRPC()

	runGateway()
}
