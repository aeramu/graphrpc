package graphrpc

import (
	"context"
	"encoding/json"
	"github.com/aeramu/graphrpc/proto"
	"github.com/graph-gophers/graphql-go"
	"google.golang.org/grpc"
)

type Server struct{
	Schema *graphql.Schema
}

func NewGraphRPCServer(schema *graphql.Schema) *grpc.Server {
	baseServer := grpc.NewServer()
	server := &Server{Schema: schema}
	proto.RegisterGraphRPCServer(baseServer, server)
	return baseServer
}

func (h *Server) Exec(ctx context.Context, request *proto.ExecRequest) (*proto.ExecResponse, error) {
	variables := map[string]interface{}{}
	if err := json.Unmarshal([]byte(request.GetVariables()), &variables); err != nil {
		return nil, err
	}

	res := h.Schema.Exec(ctx, request.GetQuery(), request.GetOperationName(), variables)
	return &proto.ExecResponse{
		Data:  string(res.Data),
		Error: nil,
	}, nil
}