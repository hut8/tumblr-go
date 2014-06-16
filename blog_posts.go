package tumblr

import (
	"github.com/google/go-querystring/query"
	"encoding/json"
	"path"
)

// Search criteria to be passed to the Posts method
type PostRequestParams struct {
	PostType   string `url:"-"`
	Id         int64  `url:"id,omitempty"`
	Tag        string `url:"tag,omitempty"`
	ReblogInfo bool   `url:"reblog_info,omitempty"`
	NotesInfo  bool   `url:"notes_info,omitempty"`
	Filter     string `url:"filter,omitempty"`
	LimitOffset
}

// TODO Use this?
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
func (blog *Blog) Posts(params PostRequestParams) (*PostCollection, error) {
	// Build URL
	url, err := blog.blogEntityURL("posts")
	if err != nil {
		return nil, err
	}

	// PostType fits in the path rather than the query string
	if params.PostType != "" {
		url.Path = path.Join(url.Path, params.PostType)
	}

	// TODO factor this out somewhere nice
	orig := url.Query()
	v, _ := query.Values(params)
	for key, val := range v {
		orig.Set(key, val[0])
	}
	url.RawQuery = orig.Encode()
	addLimitOffset(url, &params.LimitOffset)

	res, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// Decode the response without decoding the posts
	dr := &blogPostsResponse{}
	err = json.Unmarshal(*res, &dr)
	if err != nil {
		return nil, err
	}

	// Create the typed post collection
	pc, err := NewPostCollection(dr.Posts)
	return pc, err
}

type blogPostsResponse struct {
	Posts *json.RawMessage
	Total_Posts int64
}
