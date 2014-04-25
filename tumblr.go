package tumblr

import (
	"net/http"
	"net/url"
	"path"
)

func callAPI(u url.URL) (*TumblrAPIResponse, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// TODO JSON decoding
}

type TumblrAPIResponse struct {
	Meta     Meta
	Response interface{}
}

type Meta struct {
	Status int64
	Msg    string
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

type LimitOffset struct {
	Limit  int
	Offset int
}

// Request Parameter Types
type PostRequestParams struct {
	PostType   string
	Id         int64
	Tag        string
	ReblogInfo bool
	NotesInfo  bool
	Filter     string
	LimitOffset
}

func (params PostRequestParams) validatePostRequestParams() error {
	if params.Limit < 0 || params.Limit > 20 {
		return TumblrError{"Post request parameter limit out of range"}
	}
	if params.Filter != "" &&
		!(params.Filter == "html" || params.Filter == "raw") {
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

	// PostType
	if params.PostType != "" {
		url.Path = path.Join(url.Path, params.PostType)
	}

	// Id
	if params.Id != 0 {
		url.Query().Set("id", string(params.Id))
	}

	// Tag
	if params.Tag != "" {
		url.Query().Set("tag", params.Tag)
	}

	// ReblogInfo
	if params.ReblogInfo {
		url.Query().Set("reblog_info", "true")
	}

	// NotesInfo
	if params.NotesInfo {
		url.Query().Set("notes_info", "true")
	}

	addLimitOffset(url, params.LimitOffset)

	return nil
}

func addLimitOffset(url *url.URL, params LimitOffset) {
	// Limit
	if params.Limit != 0 {
		url.Query().Set("limit", string(params.Limit))
	}

	// Offset
	if params.Offset != 0 {
		url.Query().Set("offset", string(params.Offset))
	}
}
