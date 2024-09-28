package main

import (
	"connectrpc.com/connect"
	greetv1 "connectrpc_example/gen/greet/v1"
	"connectrpc_example/gen/greet/v1/greetv1connect"
	"context"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type GreetServer struct{}

func (*GreetServer) Greet(_ context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name)})
	return res, nil
}

func (*GreetServer) Mul(_ context.Context, req *connect.Request[greetv1.MulRequest]) (*connect.Response[greetv1.MulResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.MulResponse{Result: req.Msg.InputNum * 10})
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	err := http.ListenAndServe(
		"localhost:8888",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal("failed to start a server")
	}
}
