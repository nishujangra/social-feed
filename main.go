package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	mygraphql "github.com/nishujangra/social-feed/graphql"

	pb "github.com/nishujangra/social-feed/postpb"
)

var grpcClient pb.PostServiceClient

func initGRPC() {
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	grpcClient = pb.NewPostServiceClient(connection)
	if grpcClient == nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
}

func main() {
	initGRPC()

	schemaBytes, err := os.ReadFile("graphql/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(string(schemaBytes), &mygraphql.Resolver{
		PostService: grpcClient,
	})

	http.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
