package main

import (
	"chat-task/client"
	"chat-task/protos"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	envPort  = os.Getenv("PORT")
	username = flag.String("username", "", "usernam")
)

func main() {
	flag.Parse()
	if *username == "" {
		log.Fatalf("username required")
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, envPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't dial server on port %s, err=%v", envPort, err)
	}

	defer conn.Close()

	chatClient := protos.NewChatClient(conn)
	handler := client.NewCommandHandler(chatClient, *username)

	chatClientSvc := client.New(chatClient, handler, *username)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Println("[chat-client] started")
		err = chatClientSvc.Run(context.Background())
		if err != nil {
			log.Fatalf("[chat-client] terminated with error=%v", err)
		}
		wg.Done()
	}()

	go func() {
		// ignoring interupt because scan blocks
		// stdin /exit will disconnect from chat
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	}()

	wg.Wait()

	cancel()

}
