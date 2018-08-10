package main

import (
	"github.com/vikash/learning-golang/004-http-application/006-rest-api/blog"
	"github.com/zopnow/z"
	"os"
)

var handlers = z.Handlers{
	"post": blog.PostHandler{},
}

var config = z.Config{
	ServerPort: 8001,
	Database:   z.MysqlConfig{os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), "test"},
}

func main() {
	z.Service{Handlers: handlers, Config: &config}.Run()
}
