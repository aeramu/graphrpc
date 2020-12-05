package graphrpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aeramu/graphrpc/proto"
	"google.golang.org/grpc"
)

type Client struct {
	client proto.GraphRPCClient
}

func NewGraphRPCClient(cc grpc.ClientConnInterface) *Client {
	client := proto.NewGraphRPCClient(cc)
	return &Client{client: client}
}

func (c *Client) Exec(ctx context.Context, query string, variables map[string]interface{}) (string, error){
	b, err := json.Marshal(variables)
	if err != nil {
		return "", err
	}

	res, err := c.client.Exec(ctx, &proto.ExecRequest{
		Query:         query,
		OperationName: "",
		Variables:     string(b),
	})
	if err != nil {
		return "", err
	}

	return res.GetData(), errors.New(res.GetError().Message)
}
