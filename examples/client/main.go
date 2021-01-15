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
        {
            getFilm(id: "e7098dc0-b407-41e2-9a60-aca4c9638bea") {
                id
                title
				synopsis
				rating
            }
        }
    `, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("json fromat:", res)

	var data Data
	if err := json.Unmarshal([]byte(res), &data); err != nil {
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