package client

import (
	"bufio"
	"chat-task/protos"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	chatClient protos.ChatClient
	conn       protos.Chat_ConnectClient
	wg         *sync.WaitGroup
	username   string
	handler    Handler
}

func New(chatClient protos.ChatClient, handler Handler, username string) *Client {
	return &Client{
		handler:    handler,
		chatClient: chatClient,
		wg:         &sync.WaitGroup{},
		username:   username,
	}
}

func (c *Client) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	con, err := c.chatClient.Connect(ctx, &protos.ConnectRequest{
		Username: c.username,
	})
	if err != nil {
		cancel()
		return fmt.Errorf("couldn't connect to chat server, error=%v", err)
	}

	log.Printf("Logged in as: %s \n", c.username)

	c.conn = con

	c.wg.Add(2)
	go func() {
		c.handleCommands(ctx)
		cancel()
		c.wg.Done()
	}()

	go func() {
		c.incomingMessages(ctx)
		cancel()
		c.wg.Done()
	}()

	c.wg.Wait()

	return nil
}

func (c *Client) handleCommands(ctx context.Context) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("couldn't read from stdin, error=%v", err)
		}

		command := strings.TrimSpace(line)
		if command == "/exit" {
			break
		}

		err = c.handler.HandleCommand(ctx, line)
		if err != nil {
			log.Println(status.Convert(err).Message())
		}
	}

	return nil
}

func (c *Client) incomingMessages(ctx context.Context) {
	for {
		message, err := c.conn.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			return
		}
		if err != nil {
			log.Printf("couldn't receive message, error=%v", err)
			return
		}

		log.Println("---------------------------------")
		log.Println("From: ", message.GetUsername())
		log.Println("Message: ", message.GetMessage())
		log.Println("---------------------------------")
	}
}
