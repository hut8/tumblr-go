package tumblr

import (
	"net/url"
	"path"
)

type Blog struct {
	Basename string
}

type PostRequestParams struct{}

type Post struct{}

const (
	urlBase = "http://api.tumblr.com/v2/"
)
