package main

import (
	"context"
	"log"
	"net"
	"sort"
	"time"

	"google.golang.org/grpc"

	"github.com/nishujangra/social-feed/postpb"
	pb "github.com/nishujangra/social-feed/postpb"
)

type server struct {
	pb.UnimplementedPostServiceServer
}

// Mock data for posts
var mockPosts = []*postpb.Post{
	{Id: "1", AuthorId: "1", Title: "Post 1", Content: "Hello from user 1", Timestamp: time.Now().Add(-1 * time.Hour).Format(time.RFC3339)},
	{Id: "2", AuthorId: "2", Title: "Post 2", Content: "Hello from user 2", Timestamp: time.Now().Add(-2 * time.Hour).Format(time.RFC3339)},
	{Id: "3", AuthorId: "3", Title: "Post 3", Content: "Hello from user 3", Timestamp: time.Now().Add(-30 * time.Minute).Format(time.RFC3339)},
	{Id: "4", AuthorId: "4", Title: "Post 4", Content: "Hello from user 4", Timestamp: time.Now().Add(-90 * time.Minute).Format(time.RFC3339)},
	{Id: "5", AuthorId: "2", Title: "Another Post", Content: "Another from user 2", Timestamp: time.Now().Add(-10 * time.Minute).Format(time.RFC3339)},
}

// Simulated "follows" relationship
var userFollows = map[string][]string{
	"1": {"2", "3"},
	"2": {"1"},
	"3": {"4"},
	"4": {"2", "3", "5"},
	"5": {},
}

func (s *server) ListPostsByUser(ctx context.Context, req *postpb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	following := userFollows[req.UserId]

	var filteredPosts []*postpb.Post
	for _, post := range mockPosts {
		for _, follows := range following {
			if post.AuthorId == follows {
				filteredPosts = append(filteredPosts, post)
			}
		}
	}

	// sort slice, based on timestamp
	sort.Slice(filteredPosts, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339, filteredPosts[i].Timestamp)
		tj, _ := time.Parse(time.RFC3339, filteredPosts[j].Timestamp)
		return ti.After(tj)
	})

	// Limit to 20 posts in the feed
	if len(filteredPosts) > 20 {
		filteredPosts = filteredPosts[:20]
	}

	return &pb.ListPostsResponse{
		Posts: filteredPosts,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterPostServiceServer(s, &server{})
	log.Println("Post service is running on port :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
