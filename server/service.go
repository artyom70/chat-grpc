package server

import (
	"chat-task/protos"
	"context"
	"fmt"
	"os"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	clientList map[string]protos.Chat_ConnectServer
	groupList  map[string]map[string]protos.Chat_ConnectServer

	mx *sync.Mutex

	protos.UnimplementedChatServer
}

func New() *Server {
	return &Server{
		clientList: make(map[string]protos.Chat_ConnectServer),
		groupList:  make(map[string]map[string]protos.Chat_ConnectServer),
		mx:         &sync.Mutex{},
	}
}

func (s *Server) Connect(req *protos.ConnectRequest, clientConn protos.Chat_ConnectServer) error {
	if _, ok := s.clientList[req.GetUsername()]; ok {
		return fmt.Errorf("username already taken")
	}

	s.clientList[req.GetUsername()] = clientConn

	<-clientConn.Context().Done()

	s.mx.Lock()
	delete(s.clientList, req.GetUsername())
	for name, group := range s.groupList {
		delete(group, req.GetUsername())
		if len(group) == 0 {
			delete(s.groupList, name)
		}
	}
	s.mx.Unlock()

	return clientConn.Context().Err()
}

func (s *Server) JoinGroupChat(ctx context.Context, req *protos.JoinGroupChatRequest) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	s.mx.Lock()
	defer s.mx.Unlock()

	fmt.Fprintf(os.Stdout, "%q user request to join group %q \n", req.GetUsername(), req.GetGroupName())

	if _, ok := s.groupList[req.GetGroupName()]; !ok {
		return nil, fmt.Errorf("group doesn't exists")
	}

	s.groupList[req.GetGroupName()][req.GetUsername()] = s.clientList[req.GetUsername()]

	return out, nil
}

func (s *Server) LeftGroupChat(ctx context.Context, req *protos.LeftGroupChatRequest) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.groupList[req.GetGroupName()]; !ok {
		return out, fmt.Errorf("group doesn't exists")
	}

	if _, ok := s.groupList[req.GetGroupName()][req.GetUsername()]; !ok {
		return out, fmt.Errorf("you are not group member")
	}

	delete(s.groupList[req.GroupName], req.GetUsername())

	if len(s.groupList[req.GetGroupName()]) == 0 {
		delete(s.groupList, req.GetGroupName())
	}

	return out, nil
}

func (s *Server) CreateGroupChat(ctx context.Context, req *protos.CreateGroupChatRequest) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.groupList[req.GetGroupName()]; ok {
		return nil, fmt.Errorf("group already exists")
	}

	s.groupList[req.GetGroupName()] = map[string]protos.Chat_ConnectServer{}
	s.groupList[req.GetGroupName()][req.GetUsername()] = s.clientList[req.GetUsername()]

	return out, nil
}

func (s *Server) SendMessage(ctx context.Context, req *protos.SendMessageRequest) (*emptypb.Empty, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	out := new(emptypb.Empty)
	if req.GetGroupName() == "" && req.GetToUsername() == "" {
		return out, fmt.Errorf("group name or username required")
	}

	if req.GetGroupName() != "" {
		return s.sendGroupMessage(ctx, req.GetGroupName(), req.GetUsername(), req.GetMessage())
	}

	return s.sendUserMessage(ctx, req.GetUsername(), req.GetToUsername(), req.GetMessage())
}

func (s *Server) sendGroupMessage(ctx context.Context, groupName, username, message string) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	group, ok := s.groupList[groupName]
	if !ok {
		return out, fmt.Errorf("group doesn't exists")
	}

	if _, ok := group[username]; !ok {
		return out, fmt.Errorf("you are not group member")
	}

	for _, client := range group {
		err := client.Send(&protos.ReplayMessage{
			Username: username,
			Message:  message,
		})
		if err != nil {
			return out, fmt.Errorf("couldn't send message to group members")
		}
	}

	return out, nil
}

func (s *Server) sendUserMessage(ctx context.Context, username, toUsername, message string) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	if _, ok := s.clientList[toUsername]; !ok {
		return out, fmt.Errorf("user not found")
	}

	err := s.clientList[toUsername].Send(&protos.ReplayMessage{
		Username: username,
		Message:  message,
	})
	if err != nil {
		return out, fmt.Errorf("couldn't send message to user %q, err=%v", toUsername, err)
	}

	return out, nil
}

func (s *Server) ListChannels(context.Context, *emptypb.Empty) (*protos.Channels, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	channels := &protos.Channels{}
	for username := range s.clientList {
		chanelInfo := protos.ChannelInfo{
			Type: *protos.CHANNEL_TYPE_USER.Enum(),
			Name: username,
		}
		channels.Channels = append(channels.Channels, &chanelInfo)
	}

	for groupName, _ := range s.groupList {
		chanelInfo := protos.ChannelInfo{
			Type: *protos.CHANNEL_TYPE_GROUP.Enum(),
			Name: groupName,
		}
		channels.Channels = append(channels.Channels, &chanelInfo)
	}

	return channels, nil
}
