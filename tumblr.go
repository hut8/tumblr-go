package tumblr

import (
	"net/url"
	"path"
)

type Blog struct {
	BaseHostname string
}

// Request Parameter Types
type PostRequestParams struct {
	PostType *string
}

type Post struct{}

const (
	urlBase = "http://api.tumblr.com/v2/blog/"
)

// Post Types
const (
	Text   = "text"
	Quote  = "quote"
	Link   = "link"
	Answer = "answer"
	Video  = "video"
	Audio  = "audio"
	Photo  = "photo"
	Chat   = "chat"
)

// Where the request is directed to
func (Blog blog) entityURL(entityType string) (*url.URL, error) {
	url, err := url.Parse(urlBase)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, blog.BaseHostname, entityType)
	return url, nil
}

// Posts posted by a blog
func (blog Blog) Posts(params PostRequestParams) []Post {
	// TODO
	return nil
}

// Posts liked by a blog
func (blog Blog) Likes() []Post {
	// TODO
	return nil
}
