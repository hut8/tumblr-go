package tumblr

import (
	"net/url"
	"path"
	"regexp"
)

var (
	bhnUrlRegex *regexp.Regexp = regexp.MustCompile("https?://([^/]+)/")
)

type Blog struct {
	BaseHostname string
	t Tumblr
}

// Creates a new Blog object given either the proper base hostname
// or a URL containing the base hostname
func (t Tumblr) NewBlog(baseHostname string) Blog {
	if bhnUrlRegex.MatchString(baseHostname) {
		m := bhnUrlRegex.FindStringSubmatch(baseHostname)
		baseHostname = m[1]
	}
	return Blog{
		BaseHostname: baseHostname,
		t:            t,
	}
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
