package main

import (
	"context"
	"log"
	"net"
	"sort"
	"time"

	"google.golang.org/grpc"

	"github.com/nishujangra/social-feed/postpb"
)

type server struct {
	postpb.UnimplementedPostServiceServer
}

func (s *server) ListPostsByUser(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
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

	return &postpb.ListPostsResponse{
		Posts: filteredPosts,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	postpb.RegisterPostServiceServer(s, &server{})
	log.Println("Post service is running on port :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
