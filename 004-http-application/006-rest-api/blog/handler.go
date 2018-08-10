package blog

import (
	"net/http"
)

type PostHandler struct{}

func (p PostHandler) Index(r *http.Request) (interface{}, error) {
	return posts, nil
}
