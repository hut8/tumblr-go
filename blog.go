package tumblr

import (
	"net/url"
	"path"
)

type Blog struct {
	BaseHostname string
	t Tumblr
}

// Where the request is directed to
func (blog *Blog) blogEntityURL(entityType string) (*url.URL, error) {
	url, err := blog.t.apiURL()
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(
		url.Path,
		"blog",
		blog.BaseHostname,
		entityType)
	return url, nil
}
