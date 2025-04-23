package graphql

import (
	"sort"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type Post struct {
	ID        string
	AuthorID  string
	Title     string
	Content   string
	Timestamp time.Time
}

var follows = map[string][]string{
	"1": {"2", "3"}, // user 1 follows 2 and 3
	"2": {"1"},      // user 2 follows 1
	"3": {"1"},      // user 3 follows 1
}

type Resolver struct{}

type postResolver struct {
	p *Post
}

func (r *Resolver) GetTimeline(agrs struct {
	UserId string
}) ([]*postResolver, error) {
	posts := []*Post{
		{
			ID:        "1",
			AuthorID:  "1",
			Title:     "Post 1",
			Content:   "Content of Post 1",
			Timestamp: time.Now().Add(-1 * time.Hour),
		},
		{
			ID:        "2",
			AuthorID:  "2",
			Title:     "Post 2",
			Content:   "Content of Post 2",
			Timestamp: time.Now().Add(-2 * time.Hour),
		},
		{
			ID:        "3",
			AuthorID:  "3",
			Title:     "Post 3",
			Content:   "Content of Post 3",
			Timestamp: time.Now().Add(-3 * time.Hour),
		},
	}

	var filteredPosts []*postResolver
	for _, post := range posts {
		for _, following := range follows[agrs.UserId] {
			if post.AuthorID == following {
				filteredPosts = append(filteredPosts, &postResolver{p: post})
			}
		}
	}

	sort.Slice(filteredPosts, func(i, j int) bool {
		return filteredPosts[i].p.Timestamp.After(filteredPosts[j].p.Timestamp)
	})

	if len(filteredPosts) > 20 {
		filteredPosts = filteredPosts[:20]
	}

	return filteredPosts, nil
}

func (r *postResolver) ID() graphql.ID { return graphql.ID(r.p.ID) }

func (r *postResolver) AuthorID() graphql.ID { return graphql.ID(r.p.AuthorID) }

func (r *postResolver) Title() string { return r.p.Title }

func (r *postResolver) Content() string { return r.p.Content }

func (r *postResolver) Timestamp() string {
	return r.p.Timestamp.Format(time.RFC3339)
}
