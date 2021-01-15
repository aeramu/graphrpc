package graphrpc

import (
	"context"
)

type Request struct {
	Client        *Client
	Query         string
	OperationName string
	Variables     map[string]interface{}
}

func (r Request) Exec(ctx context.Context) (*Response, error) {
	res, err := r.Client.Exec(ctx, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r Request) ExecQuery(ctx context.Context, query string) (*Response, error) {
	r.Query = query

	res, err := r.Client.Exec(ctx, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r Request) ExecWithVariables(ctx context.Context, variables map[string]interface{}) (*Response, error) {
	r.Variables = variables

	res, err := r.Client.Exec(ctx, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r Request) WithQuery(query string) Request {
	r.Query = query
	return r
}

func (r Request) WithVariables(variables map[string]interface{}) Request {
	r.Variables = variables
	return r
}

func (r Request) WithOperationName(operationName string) Request {
	r.OperationName = operationName
	return r
}
