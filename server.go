package graphrpc

import (
	"context"
	"encoding/json"
	"github.com/aeramu/graphrpc/proto"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
	"google.golang.org/grpc"
	"log"
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
		Errors: parseErrors(res.Errors),
	}, nil
}

func parseErrors(errs []*errors.QueryError) (protoErrors []*proto.Error) {
	for _, err := range errs {
		protoErrors = append(protoErrors, &proto.Error{
			Message: err.Message,
			Locations: parseLocations(err.Locations),
			Path: parsePath(err.Path),
		})
	}
	return
}

func parseLocations(locations []errors.Location) (protoLocations []*proto.Location) {
	for _, location := range locations {
		protoLocations = append(protoLocations, &proto.Location{
			Line:   int32(location.Line),
			Column: int32(location.Column),
		})
	}
	return
}

func parsePath(path []interface{}) (protoPath []string) {
	for _, p := range path {
		b, err := json.Marshal(p)
		if err != nil {
			log.Println("failed marshal graphql error path:", err)
			return nil
		}
		protoPath = append(protoPath, string(b))
	}
	return
}
