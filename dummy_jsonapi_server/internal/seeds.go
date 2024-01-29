package internal

import (
	"time"
)

func seedData() *Blog {
	return &Blog{
		ID:        1,
		Title:     "My First Blog",
		Posts:     seedPosts(),
		CreatedAt: time.Now(),
		ViewCount: 100,
	}
}

func seedPosts() []*Post {
	return []*Post{
		{
			ID:       1,
			BlogID:   1,
			Title:    "My First Post",
			Body:     "This is the body of my first post",
			Comments: seedComments(0, 1),
		},
		{
			ID:       2,
			BlogID:   1,
			Title:    "My Second Post",
			Body:     "This is the body of my second post",
			Comments: seedComments(2, 2),
		},
	}
}

func seedComments(id int, p_id int) []*Comment {
	return []*Comment{
		{
			ID:     id + 1,
			PostID: p_id,
			Body:   "This is the first comment",
		},
		{
			ID:     id + 2,
			PostID: p_id,
			Body:   "This is the second comment",
		},
	}
}
