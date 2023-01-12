package client

import "context"

type Handler interface {
	HandleCommand(ctx context.Context, data string) error
}
