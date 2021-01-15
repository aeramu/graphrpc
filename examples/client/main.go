package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aeramu/graphrpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
	client := graphrpc.NewGraphRPCClient(conn)

	res, err := client.Exec(context.Background(), `
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
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("graphql errors:", res.Errors)
	fmt.Println("json fromat:", res.Data)
	var data Data
	if err := json.Unmarshal([]byte(res.Data), &data); err != nil {
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