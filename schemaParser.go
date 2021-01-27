package graphrpc

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/graph-gophers/graphql-go"
)

func ParseSchemaFromFile(dir string, resolver interface{}) (*graphql.Schema, error) {
	f, err := os.Open(dir)
	if err != nil {
		log.Println("Failed open graphql schema file")
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Failed read graphql schema")
		return nil, err
	}

	schema, err := graphql.ParseSchema(string(b), resolver)
	if err != nil {
		log.Println("Failed parse graphql schema with resolver")
		return nil, err
	}

	return schema, nil
}
