package server

import (
	"context"
	"log"
	"net"

	pb "github.com/rafaelmbcosta/golang-blog-experiment/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))

	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect to client: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	// end of db configuration

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
