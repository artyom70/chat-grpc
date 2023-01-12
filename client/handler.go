package client

import (
	"chat-task/protos"
	"context"
	"log"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CommandHandler struct {
	chatClient protos.ChatClient
	username   string
}

func NewCommandHandler(chatClient protos.ChatClient, username string) *CommandHandler {
	return &CommandHandler{
		chatClient: chatClient,
		username:   username,
	}
}

func (h *CommandHandler) HandleCommand(ctx context.Context, data string) error {
	commandArr := strings.Fields(data)
	if len(commandArr) == 0 {
		return nil
	}

	command := commandArr[0]

	switch command {
	case "/jg":
		groupName, err := h.decodeJoinGroupCommand(data)
		if err != nil {
			return err
		}
		_, err = h.chatClient.JoinGroupChat(ctx, &protos.JoinGroupChatRequest{
			Username:  h.username,
			GroupName: groupName,
		})
		if err != nil {
			return err
		}
	case "/lg":
		groupName, err := h.decodeLeftGroupChatCommand(data)
		if err != nil {
			return err
		}
		_, err = h.chatClient.LeftGroupChat(ctx, &protos.LeftGroupChatRequest{
			Username:  h.username,
			GroupName: groupName,
		})
		if err != nil {
			return err
		}
	case "/cg":
		groupName, err := h.decodeCreateGroupChatCommand(data)
		if err != nil {
			return err
		}
		_, err = h.chatClient.CreateGroupChat(ctx, &protos.CreateGroupChatRequest{
			Username:  h.username,
			GroupName: groupName,
		})
		if err != nil {
			return err
		}
	case "/smg":
		groupName, message, err := h.decodeSendMessageGroupCommand(data)
		if err != nil {
			return err
		}

		_, err = h.chatClient.SendMessage(ctx, &protos.SendMessageRequest{
			GroupName: groupName,
			Username:  h.username,
			Message:   message,
		})
		if err != nil {
			return err
		}
	case "/sm":
		toUsername, message, err := h.decodeSendMessageCommand(data)
		if err != nil {
			return err
		}
		_, err = h.chatClient.SendMessage(ctx, &protos.SendMessageRequest{
			Username:   h.username,
			Message:    message,
			ToUsername: toUsername,
		})
		if err != nil {
			return err
		}
	case "/lc":
		var out strings.Builder

		in := new(emptypb.Empty)
		channels, err := h.chatClient.ListChannels(ctx, in)
		if err != nil {
			return err
		}
		for _, channel := range channels.GetChannels() {
			out.WriteString("\n")
			out.WriteString("Type: ")
			out.WriteString(converChannelType(channel.GetType()))
			out.WriteString("\n")
			out.WriteString("Name: ")
			out.WriteString(channel.GetName())

		}
		log.Println(out.String())

	default:
		return ErrInvalidCommand
	}

	return nil
}

func converChannelType(channelType protos.CHANNEL_TYPE) string {
	switch channelType {
	case protos.CHANNEL_TYPE_GROUP:
		return "group"
	case protos.CHANNEL_TYPE_USER:
		return "user"
	default:
		return "unknown"
	}
}

func (h *CommandHandler) decodeJoinGroupCommand(data string) (string, error) {
	command := strings.Split(data, " ")
	if len(command) < 2 {
		return "", ErrInvalidCommand
	}

	return strings.TrimSpace(command[1]), nil
}

func (h *CommandHandler) decodeLeftGroupChatCommand(data string) (string, error) {
	command := strings.Split(data, " ")
	if len(command) < 2 {
		return "", ErrInvalidCommand
	}

	return strings.TrimSpace(command[1]), nil
}

func (h *CommandHandler) decodeCreateGroupChatCommand(data string) (string, error) {
	command := strings.Split(data, " ")
	if len(command) < 2 {
		return "", ErrInvalidCommand
	}

	return strings.TrimSpace(command[1]), nil
}

func (h *CommandHandler) decodeSendMessageGroupCommand(data string) (string, string, error) {
	command := strings.Split(data, " ")
	if len(command) < 2 {
		return "", "", ErrInvalidCommand
	}

	message := strings.TrimSpace(data[len(strings.TrimSpace(command[0]))+len(strings.TrimSpace(command[1]))+1:])

	return strings.TrimSpace(command[1]), strings.TrimSpace(message), nil
}

func (h *CommandHandler) decodeSendMessageCommand(data string) (string, string, error) {
	command := strings.Split(data, " ")
	if len(command) < 2 {
		return "", "", ErrInvalidCommand
	}

	message := strings.TrimSpace(data[len(strings.TrimSpace(command[0]))+len(strings.TrimSpace(command[1]))+1:])

	return strings.TrimSpace(command[1]), strings.TrimSpace(message), nil
}
