package tumblr

import (
	"path"
)

type PostRequestParams struct {
	PostType   string `url:"-"`
	Id         int64  `url:"id,omitempty"`
	Tag        string `url:"tag,omitempty"`
	ReblogInfo bool   `url:"reblog_info,omitempty"`
	NotesInfo  bool   `url:"notes_info,omitempty"`
	Filter     string `url:"filter,omitempty"`
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
func (blog *Blog) Posts(params PostRequestParams) ([]Post, error) {
	// Build URL
	url, err := blog.blogEntityURL("posts")
	if err != nil {
		return nil, err
	}

	// TODO use go-querystring here!!!!!!!!!!!
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

	data, err := callAPI(url)

	var posts []Post
	// TODO Deserialize results

	return posts, nil
}
