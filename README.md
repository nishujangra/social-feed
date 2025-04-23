# 📘 Social Feed GraphQL API

A simple GraphQL API built in Go to simulate a user timeline. It returns the 20 most recent posts from users that a given user follows — sorted in reverse chronological order.

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

2. **Install dependencies**  
   Make sure Go is installed (v1.18+ recommended)

   ```bash
   go mod tidy
   ```

3. **Ensure schema file exists**  
   Your `schema.graphql` should be at the root and contain:

   ```graphql
   type Query {
     getTimeline(userId: ID!): [Post]
   }

   type Post {
     id: ID!
     authorId: ID!
     title: String!
     content: String!
     createdAt: String!
   }
   ```

---

## ▶️ How to Run the Service

```bash
go run main.go
```

This starts the server on:

```
http://localhost:8080/graphql
```

---

## 🧪 Sample GraphQL Query (via cURL)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ getTimeline(userId: \"1\") { id authorId title content createdAt } }"}'
```

---

## 🧠 Description of the Approach

- **GraphQL Server** is built using [`graph-gophers/graphql-go`](https://github.com/graph-gophers/graphql-go)
- A `getTimeline(userId: ID!)` resolver checks which users the given `userId` follows
- Posts by followed users are filtered from an in-memory slice
- The posts are sorted by their `timestamp` field in reverse chronological order
- The resolver returns the top 20 (or fewer) posts

This structure allows for easy future enhancements like:
- Using PostgreSQL/MongoDB
- Adding pagination, likes, comments
- User authentication and role-based feeds

---

## 📄 License

MIT License – see the [LICENSE](LICENSE.md) file for details.