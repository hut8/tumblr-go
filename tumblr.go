package tumblr

import (
	"fmt"
	"github.com/kr/pretty"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func callAPI(u *url.URL) (interface{}, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	json.Unmarshal(data, &raw)

	// How the crap am I supposed to do this?!
	meta := Meta{
		Status: int64((raw["meta"]).(map[string]interface{})["status"].(float64)),
		Msg: (raw["meta"]).(map[string]interface{})["msg"].(string),
	}

	if meta.Status != 200 {
		err = fmt.Errorf("tumblr API responded with HTTP status %d: %s",
			meta.Status,
			meta.Msg)
		return nil, err
	}

	res := raw["response"]

	fmt.Printf("response for %v:\n%v\n", u, pretty.Formatter(res))

	return res, nil
}

type TumblrAPIResponse struct {
	Meta     Meta
	Response interface{}
}

type Meta struct {
	Status int64
	Msg    string
}

type APICredentials struct {
	Key string
	Secret string
}

type Blog struct {
	BaseHostname string
	Credentials APICredentials
}

type Post struct{}

const (
	urlBaseBlog = "http://api.tumblr.com/v2/blog/"
	urlBaseUser = "http://api.tumblr.com/v2/user/"
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
func (blog Blog) blogEntityURL(entityType string) (*url.URL, error) {
	url, err := url.Parse(urlBaseBlog)
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
	url, err := blog.blogEntityURL("posts")
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
