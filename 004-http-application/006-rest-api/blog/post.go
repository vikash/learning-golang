package blog

import "time"

type Post struct {
	Id          int
	Title       string
	Body        string
	PublishTime time.Time
}

var posts = []Post{
	{1, "Hello World!", "This is our first post.", time.Now()},
	{2, "Another World!", "This is our second post.", time.Now()},
}

