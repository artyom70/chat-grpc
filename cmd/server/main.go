package main

import (
	"chat-task/protos"
	"chat-task/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	envPort = os.Getenv("PORT")
)

func main() {
	listener, err := net.Listen("tcp", envPort)
	if err != nil {
		log.Fatalf("couldn't listent on port, %s ", envPort)
	}

	srv := grpc.NewServer()
	cs := server.New()
	protos.RegisterChatServer(srv, cs)

	log.Printf("[chat-server] started listening on port %s", envPort)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to start gRPC server, err=%v ", err)
	}

}
