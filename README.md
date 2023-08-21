# echo-graphql

```graphql
mutation createPost {
  createPost(input: {id: "1", title: "Title", content: "Content"}) {
    id
    title
    content
  }
}
```

```graphql
mutation createPost {
  updatePost(input: {id: "2", title: "TitleChange", content: "ContentChange"}) {
    id
    title
    content
  }
}
```

```graphql
mutation deletePost {
  deletePost(input: {id: "3"}) {
    id
    title
    content
  }
}
```

```graphql
query getPosts {
  posts {
    id
    title
    content
  }
}
```

```graphql
query getPosts {
  post(id: "1") {
    id
    title
    content
  }
}
```
