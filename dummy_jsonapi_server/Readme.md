# To start server
- Run `go run cmd/server.go`
- Visit `curl --location 'http://localhost:9999/' --header 'Accept: application/vnd.api+json'`
  
Received Output:

```{
    "data": {
        "type": "blogs",
        "id": "1",
        "attributes": {
            "created_at": 1706543465,
            "current_post_id": 0,
            "title": "My First Blog",
            "view_count": 100
        },
        "relationships": {
            "current_post": {
                "data": null
            },
            "posts": {
                "data": [
                    {
                        "type": "posts",
                        "id": "1"
                    },
                    {
                        "type": "posts",
                        "id": "2"
                    }
                ],
                "links": {
                    "related": "https://example.com/blogs/1/posts"
                },
                "meta": {
                    "detail": "posts meta information"
                }
            }
        },
        "links": {
            "self": "https://example.com/blogs/1"
        },
        "meta": {
            "detail": "extra details regarding the blog"
        }
    },
    "included": [
        {
            "type": "comments",
            "id": "1",
            "attributes": {
                "body": "This is the first comment",
                "post_id": 1
            }
        },
        {
            "type": "comments",
            "id": "2",
            "attributes": {
                "body": "This is the second comment",
                "post_id": 1
            }
        },
        {
            "type": "comments",
            "id": "3",
            "attributes": {
                "body": "This is the first comment",
                "post_id": 2
            }
        },
        {
            "type": "comments",
            "id": "4",
            "attributes": {
                "body": "This is the second comment",
                "post_id": 2
            }
        },
        {
            "type": "posts",
            "id": "1",
            "attributes": {
                "blog_id": 1,
                "body": "This is the body of my first post",
                "title": "My First Post"
            },
            "relationships": {
                "comments": {
                    "data": [
                        {
                            "type": "comments",
                            "id": "1"
                        },
                        {
                            "type": "comments",
                            "id": "2"
                        }
                    ]
                }
            }
        },
        {
            "type": "posts",
            "id": "2",
            "attributes": {
                "blog_id": 1,
                "body": "This is the body of my second post",
                "title": "My Second Post"
            },
            "relationships": {
                "comments": {
                    "data": [
                        {
                            "type": "comments",
                            "id": "3"
                        },
                        {
                            "type": "comments",
                            "id": "4"
                        }
                    ]
                }
            }
        }
    ]
}
```




- Make a post call
```curl --location --request GET 'http://localhost:9999/' \
  --header 'Accept: application/vnd.api+json' \
  --header 'Content-Type: application/json' \
  --data '{
  "id": 2,
  "blog_id": "1",
  "title": "asdfasdfasdfafds",
  "body": "adfasdfasdfasdfsdasdafsadf",
  "comments": []
  }
```
Received: `data is not a jsonapi representation of '*internal.Post'`. Look more into this.
