package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/aeramu/graphrpc"
	"github.com/graph-gophers/graphql-go"
)

func main() {
	f, err := os.Open("schema.graphql")
	if err != nil {
		log.Println("failed open graphql schema file")
		return
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("failed read graphql schema")
		return
	}
	schema := graphql.MustParseSchema(string(b), &Resolver{})
	server := graphrpc.NewGraphRPCServer(schema)

	listener, _ := net.Listen("tcp", ":8000")
	log.Println("server started at :8000")
	log.Fatal(server.Serve(listener))
}