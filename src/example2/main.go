package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"pb"
	"services"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "port of all services")
)

func runGRPC(l net.Listener) {
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &services.EchoService{})

	s.Serve(l)
}

func runGateway(l net.Listener) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEchoServiceHandlerFromEndpoint(context.TODO(), mux, fmt.Sprintf(":%d", *port), opts)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
		}),
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}

	return
}

func main() {
	flag.Parse()

	// Create the main listener.
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpl := m.Match(cmux.HTTP1Fast())

	go runGRPC(grpcl)
	go runGateway(httpl)

	fmt.Println("grpc server started.")
	fmt.Println("http server started.")
	fmt.Println("Server listening on port", *port)

	// Start cmux serving.
	if err := m.Serve(); err != nil {
		panic(err)
	}
}
