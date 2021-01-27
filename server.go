package graphrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aeramu/graphrpc/proto"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Schema *graphql.Schema
}

func NewGraphRPCServer(schema *graphql.Schema) *grpc.Server {
	baseServer := grpc.NewServer()
	server := &Server{Schema: schema}
	proto.RegisterGraphRPCServer(baseServer, server)
	reflection.Register(baseServer)
	return baseServer
}

func (h *Server) Exec(ctx context.Context, request *proto.ExecRequest) (*proto.ExecResponse, error) {
	variables := map[string]interface{}{}
	if request.GetVariables() != nil {
		if err := json.Unmarshal(request.GetVariables(), &variables); err != nil {
			log.Println("Failed unmarshal JSON variables:", err)
			return nil, err
		}
	}

	res := h.Schema.Exec(ctx, request.GetQuery(), request.GetOperationName(), variables)
	return &proto.ExecResponse{
		Data:   res.Data,
		Errors: parseErrors(res.Errors),
	}, nil
}

func parseErrors(errs []*errors.QueryError) (protoErrors []*proto.Error) {
	for _, err := range errs {
		protoErrors = append(protoErrors, &proto.Error{
			Message:   err.Message,
			Locations: parseLocations(err.Locations),
			Path:      parsePath(err.Path),
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
		protoPath = append(protoPath, fmt.Sprintf("%v", p))
	}
	return
}
