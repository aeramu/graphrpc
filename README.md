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
    "github.com/aeramu/graphrpc"
    "google.golang.org/grpc"
)

func main() {
    conn, _ := grpc.Dial("localhost:8000", grpc.WithInsecure())
    client := graphrpc.NewGrahRPCClient(conn)

    res, err := client.Exec(context.Background(), `
        {
            film {
                id
                title
                actor {
                    id
                    name
                }
            }
        }
    `, nil)    	
}

```

## TODO
- Testing
- GraphQL Error (Error handling not implemented yet)
- Object Mapping (Response Consume function)
- Graphql Variables (Is map[string]interface{} good?)
- Exec Builder for client
- Add support to different GraphQL library(?) 
(or maybe make another own)