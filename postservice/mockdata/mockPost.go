package mockdata

import (
	"time"

	"github.com/nishujangra/social-feed/postpb"
)

// Mock data for posts
var MockPosts = []*postpb.Post{
	{Id: "1", AuthorId: "1", Title: "Go Routines", Content: "Concurrency made easy!", Timestamp: time.Now().Add(-1 * time.Hour).Format(time.RFC3339)},
	{Id: "2", AuthorId: "2", Title: "Microservices", Content: "Why I prefer gRPC over REST", Timestamp: time.Now().Add(-2 * time.Hour).Format(time.RFC3339)},
	{Id: "3", AuthorId: "3", Title: "Travel Vlog", Content: "Just came back from the mountains 🏔️", Timestamp: time.Now().Add(-30 * time.Minute).Format(time.RFC3339)},
	{Id: "4", AuthorId: "4", Title: "Book Review", Content: "Deep Work is a must-read for devs", Timestamp: time.Now().Add(-90 * time.Minute).Format(time.RFC3339)},
	{Id: "5", AuthorId: "2", Title: "Weekend Plans", Content: "Thinking of building a CLI tool", Timestamp: time.Now().Add(-10 * time.Minute).Format(time.RFC3339)},
	{Id: "6", AuthorId: "5", Title: "Photography", Content: "Captured a stunning sunset 🌇", Timestamp: time.Now().Add(-25 * time.Minute).Format(time.RFC3339)},
	{Id: "7", AuthorId: "3", Title: "Coffee ☕️", Content: "Brewed the perfect cup today", Timestamp: time.Now().Add(-5 * time.Minute).Format(time.RFC3339)},
}

// Simulated "follows" relationship
var UserFollows = map[string][]string{
	"1": {"2", "3"},      // User 1 follows 2 and 3
	"2": {"1", "3", "5"}, // User 2 follows 1, 3, and 5
	"3": {"4"},           // User 3 follows 4
	"4": {"2", "3", "5"}, // User 4 follows 2, 3, and 5
	"5": {"1", "4"},      // User 5 follows 1 and 4
}
