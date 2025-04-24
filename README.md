# 📘 Social Feed - GraphQL + gRPC

This project is a simple social feed service using:
- GraphQL API (via `graph-gophers/graphql-go`)
- gRPC service for posts data (mocked for now)

---

## 🚀 Features

- 🔍 `getTimeline(userId: ID!)` GraphQL query
- 🕒 Returns 20 most recent posts from followed users
- 🔃 Posts sorted by `timestamp` (latest first)
- 🧠 Uses in-memory mock data (can easily plug in a DB)

---

## ⚙️ Setup Instructions

1. **Clone the repository**  
   ```bash
   git clone https://github.com/nishujangra/social-feed.git
   cd social-feed
   ```

2. **Install Go dependencies**  
   Make sure Go is installed (v1.18+ recommended)

   ```bash
   go mod tidy
   ```

3. **Install protoc and Go plugins (if not already installed)**
    In Ubuntu/Debian

    ```bash
      sudo apt install -y protobuf-compiler
      go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
      go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

    Ensure `$GOPATH/bin` is in your `PATH`:

    ```bash
      export PATH="$PATH:$(go env GOPATH)/bin"
    ```
    for single terminal session.

4. **Generate gRPC code from proto file**

    ```bash
      protoc --go_out=. --go-grpc_out=. /proto/post.proto
    ```


5. **Run gRPC Post Service**
   ```bash
   cd postservice
   go run server.go
   ```

   The gRPC service runs on: `localhost:50051` (default gRPC port)

6. **Run GraphQL Server**
   In another terminal:
   ```bash
   go run main.go
   ```

   The GraphQL server runs on: `http://localhost:8080/graphql` (HTTP)

---

## 📁 Project Structure

```
.
├── graphql/                 # GraphQL schema and resolvers
│   ├── resolver.go
│   └── schema.graphql
├── proto/
│      └── post.proto
├── postservice/            # Simulated gRPC Post Service
│   ├── mockdata/
│       └── mockPost.go
│   └── server.go
├── main.go                 # Starts the GraphQL server
├── go.mod
├── .gitignore
└── README.md
```

---

## 🚀 How to Run the GraphQL Service

```bash
go run main.go
```

By default, the GraphQL server runs on:
```
http://localhost:8080/graphql
```

You can test it using GraphiQL, Postman, or `curl`.

---

## ▶️ How to Run the Post Services

```bash
go run postservice/server.go
```
---

## 🧪 Mock Data Configuration

The service uses simulated data for testing:

### Posts Data (`postservice/mockdata/mockPost.go`)
```go
var MockPosts = []*postpb.Post{
  {Id: "1", AuthorId: "1", Title: "Go Routines", ...},
  {Id: "2", AuthorId: "2", Title: "Microservices", ...},
  // ... 7 total mock posts
}
```

### Follow Relationships
```go
var UserFollows = map[string][]string{
  "1": {"2", "3"},      // User 1 follows 2 and 3
  "2": {"1", "3", "5"}, // User 2 follows 1, 3 and 5
  // ... more relationships
}
```

### Key Details:
- **Pre-configured Users**: IDs 1-5
- **Sample Content**: Tech posts, travel blogs, and personal updates
- **Timestamps**: Posts span from 5-120 minutes old (relative to server start)

### Customization:
Edit the mock data by modifying:
```
postservice/mockdata/mockPost.go
```
Changes take effect when you restart the gRPC server.

---

## 🧪 Sample GraphQL Query (via cURL)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ getTimeline(userId: \"1\") { id authorId title content timestamp } }"}'
```

---


## 📦 Sample GraphQL Query

```graphql
query {
  getTimeline(userId: "1") {
    id
    authorId
    title
    content
    timestamp
  }
}
```

---

## 🧠 Description of the Approach

- **GraphQL Server** is built using [`graph-gophers/graphql-go`](https://github.com/graph-gophers/graphql-go)
- The `getTimeline(userId)` GraphQL resolver communicates with a gRPC Post Service.
- The gRPC service:
  - Mocks users and their posts.
  - Maintains a follow map: which users are followed by whom.
  - Filters posts where the author is followed by the given user.
  - Combines and sorts all followed users' posts by timestamp.
  - Returns the **latest 20 posts**.

This structure allows for easy future enhancements like:
- Using PostgreSQL/MongoDB
- Adding pagination, likes, comments
- User authentication and role-based feeds

---

## 📄 License

MIT License – see the [LICENSE](LICENSE.md) file for details.