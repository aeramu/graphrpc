package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aeramu/graphrpc"
	"google.golang.org/grpc"
)

const query = `
	query($id: ID!){
		getFilm(id: $id) {
			id
			title
			synopsis
			rating
		}
	}`

func main() {
	conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
	client := graphrpc.NewGraphRPCClient(conn)

	res, err := client.
		WithQuery(query).
		WithVariables(map[string]interface{}{
			"id": "e7098dc0-b407-41e2-9a60-aca4c9638bea",
		}).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	var data Data
	if err := res.Consume(&data); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(data)
}

type Data struct {
	GetFilm struct {
		ID       string
		Title    string
		Synopsis string
		Rating   float64
	}
}
