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
	return &Client{
		client: client,
	}
}

func (c *Client) Exec(ctx context.Context, req Request) (*Response, error) {
	if c == nil {
		return nil, errors.New("client is nil")
	}

	if c.client == nil {
		return nil, errors.New("grpc client is nil")
	}

	variables, err := json.Marshal(req.Variables)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Exec(ctx, &proto.ExecRequest{
		Query:         req.Query,
		OperationName: req.OperationName,
		Variables:     variables,
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Data:   res.GetData(),
		Errors: res.GetErrors(),
	}, nil
}

func (c *Client) ExecQuery(ctx context.Context, query string) (*Response, error) {
	req := Request{
		Query:         query,
		OperationName: "",
		Variables:     nil,
	}

	res, err := c.Exec(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) WithQuery(query string) Request {
	return Request{
		Client: c,
		Query:  query,
	}
}

func (c *Client) WithVariables(variables map[string]interface{}) Request {
	return Request{
		Client:    c,
		Variables: variables,
	}
}

func (c *Client) WithOperationName(operationName string) Request {
	return Request{
		Client:        c,
		OperationName: operationName,
	}
}
