package graphrpc

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *Client) Exec(ctx context.Context, query string, variables map[string]interface{}) *Response {
	b, err := json.Marshal(variables)
	if err != nil {
		return nil
	}

	res, err := c.client.Exec(ctx, &proto.ExecRequest{
		Query:         query,
		OperationName: "",
		Variables:     b,
	})
	if err != nil {
		return &Response{
			Data:          nil,
			GrpcError:     err,
			GraphqlErrors: nil,
		}
	}

	return &Response{
		Data:          res.GetData(),
		GraphqlErrors: res.GetErrors(),
		GrpcError:     nil,
	}
}

type Response struct{
	Data          []byte
	GrpcError     error
	GraphqlErrors []*proto.Error
}

func (r Response) Consume(v interface{}) error {
	if r.GrpcError != nil {
		return r.GrpcError
	}
	if len(r.GraphqlErrors) > 0 {
		return fmt.Errorf("%v", r.GraphqlErrors)
	}
	if err := json.Unmarshal(r.Data, &v); err != nil {
		return err
	}
	return nil
}