package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	mygraphql "github.com/nishujangra/social-feed/graphql"
)

func main() {
	schemaBytes, err := os.ReadFile("graphql/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(string(schemaBytes), &mygraphql.Resolver{})

	http.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
