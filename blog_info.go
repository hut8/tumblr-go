package tumblr

import (
	"encoding/json"
)

// Blog Info
func (blog Blog) Info() (*BlogInfo, error) {
	url, err := blog.blogEntityURL("info")
	if err != nil {
		return nil, err
	}

	res, err := callAPI(url)
	if err != nil {
		return nil, err
	}
	m, _ := res.MarshalJSON()

	i := &blogInfoResponse{}
	err = json.Unmarshal(*res, i)
	if err != nil {
		return nil, err
	}

	return &i.Blog, nil
}

type blogInfoResponse struct {
	Blog BlogInfo
}

// Type returned by blog.Info()
type BlogInfo struct {
	Title       string
	Posts       int64
	Name        string
	Updated     int64
	Description string
	Ask         bool
	AskAnon     bool
	Likes       int64
}
