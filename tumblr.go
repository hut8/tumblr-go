package tumblr

import (
	"net/url"
	"path"
)

// Errors
type TumblrError struct {
	Message string
}

func (err TumblrError) Error() string {
	return err.Message
}

type Blog struct {
	BaseHostname string
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
func (blog Blog) entityURL(entityType string) (*url.URL, error) {
	url, err := url.Parse(urlBase)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, blog.BaseHostname, entityType)
	return url, nil
}

// Request Parameter Types
type PostRequestParams struct {
	PostType   *string
	Id         *int64
	Tag        *string
	Limit      int
	Offset     int
	ReblogInfo bool
	NotesInfo  bool
	Filter     *string
}

func (params PostRequestParams) validatePostRequestParams() error {
	if params.Limit < 1 || params.Limit > 20 {
		return TumblrError{"Post request parameter limit out of range"}
	}
	if params.Filter != nil &&
		!(*params.Filter == "html" || *params.Filter == "raw") {
		return TumblrError{`Filter, if specified, must be either "html" or "raw"`}
	}
	return nil
}

// Posts posted by a blog
func (blog Blog) Posts(params PostRequestParams) []Post {
	// Build URL
	url, err := blog.entityURL("posts")
	if err != nil {
		return nil
	}
	if params.PostType != nil {
		url.Path = path.Join(url.Path, *params.PostType)
	}

	// Query string parameters
	if params.ReblogInfo {
		url.Query().Set("reblog_info", "true")
	}
	if params.NotesInfo {
		url.Query().Set("notes_info", "true")
	}

	return nil
}

// Posts liked by a blog
func (blog Blog) Likes() []Post {
	// TODO
	return nil
}
