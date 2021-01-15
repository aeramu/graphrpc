package graphrpc

import (
	"context"
	"encoding/json"
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

func (c *Client) Exec(ctx context.Context, query string, variables map[string]interface{}) (*Response, error){
	b, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Exec(ctx, &proto.ExecRequest{
		Query:         query,
		OperationName: "",
		Variables:     string(b),
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Data:   res.GetData(),
		Errors: res.GetErrors(),
	}, nil
}

type Response struct{
	Data string
	Errors []*proto.Error
}
