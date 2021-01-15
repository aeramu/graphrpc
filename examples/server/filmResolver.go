package main

import "github.com/graph-gophers/graphql-go"

type filmResolver struct {}

func (r *filmResolver) ID() graphql.ID {
	return "e7098dc0-b407-41e2-9a60-aca4c9638bea"
}

func (r *filmResolver) Title() string {
	return "New Film"
}

func (r *filmResolver) Synopsis() string {
	return "lorem ipsum"
}

func (r *filmResolver) Rating() float64 {
	return 4.8
}
