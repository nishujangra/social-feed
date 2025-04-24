package main

import (
	"time"

	"github.com/nishujangra/social-feed/postpb"
)

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
