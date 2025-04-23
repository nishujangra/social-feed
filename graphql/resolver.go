package graphql

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	pb "github.com/nishujangra/social-feed/postpb"
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

type Resolver struct {
	PostService pb.PostServiceClient
}

type postResolver struct {
	p *Post
}

func (r *Resolver) GetTimeline(agrs struct {
	UserId string
}) ([]*postResolver, error) {
	ctx := context.Background()
	response, err := r.PostService.ListPostsByUser(ctx, &pb.ListPostsRequest{
		UserId: agrs.UserId,
	})
	if err != nil {
		return nil, err
	}

	var posts []*postResolver
	for _, post := range response.Posts {
		posts = append(posts, &postResolver{
			p: &Post{
				ID:        post.Id,
				AuthorID:  post.AuthorId,
				Title:     post.Title,
				Content:   post.Content,
				Timestamp: time.Now(), // Assuming the timestamp is set to now for simplicity
			},
		})
	}
	return posts, nil
}

func (r *postResolver) ID() graphql.ID { return graphql.ID(r.p.ID) }

func (r *postResolver) AuthorID() graphql.ID { return graphql.ID(r.p.AuthorID) }

func (r *postResolver) Title() string { return r.p.Title }

func (r *postResolver) Content() string { return r.p.Content }

func (r *postResolver) Timestamp() string {
	return r.p.Timestamp.Format(time.RFC3339)
}
