package main

import (
	"context"
	"fmt"
	"github.com/aeramu/graphrpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
	client := graphrpc.NewGraphRPCClient(conn)

	res := client.Exec(context.Background(), `
        query($id: ID!){
            getFilm(id: $id) {
                id
                title
				synopsis
				rating
            }
        }
    `, map[string]interface{}{
		"id": "e7098dc0-b407-41e2-9a60-aca4c9638bea",
	})

	var data Data
	if err := res.Consume(&data); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("go struct:", data)
}

type Data struct {
	GetFilm struct {
		ID string
		Title string
		Synopsis string
		Rating float64
	}
}