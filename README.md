# GraphRPC
**GraphQL over gRPC**. 

GraphQL has made 
query and mutating data easier for front-end.
How about using GraphQL to communicate between microservices?
And lately, gRPC has been choosen over REST to communicate
between microservices in backend because of the speed.

Here,
**GraphRPC made the transport of the GraphQL query to gRPC**
for easy communicate between microservices without 
neglecting the speed

Note: GraphRPC intended to extend 
[graph-gophers](github.com/graph-gophers/graphql-go)
handler over gRPC. GraphRPC need GraphQL resolver 
implementation from it.
## Usages
Complete example see ```examples``` folder
### Get Started
```bash
go get -u github.com/aeramu/graphrpc
```
### Server Example
```go
package main

import (
    "net"

    "github.com/aeramu/graphrpc"
    "github.com/graph-gophers/graphql-go"
)

func main() {
    schema := graphql.MustParseSchema(schemaString, &resolver{})
    server := graphrpc.NewGraphRPCServer(schema)

    listener, _ := net.Listen("tcp", ":8000")
    server.Serve(listener)
}

```
### Client Example
```go
package main

import (
    "context"
    "log"

    "github.com/aeramu/graphrpc"
    "google.golang.org/grpc"
)

func main() {
    conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
    client := graphrpc.NewGraphRPCClient(conn)

    res, err := client.
    	WithQuery(
    	`query($id: ID!){
            getFilm(id: $id) {
                id
                title
                synopsis
                rating
            }
        }`).
        ExecWithVariables(context.Background(), map[string]interface{}{
            "id": "e7098dc0-b407-41e2-9a60-aca4c9638bea",
        })  
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

```

## TODO
- Testing
- Is API good enough? (Exec Request Builder)
- Is it concurrency safe?
- Add support to different GraphQL library(?) 
(or maybe make another own)