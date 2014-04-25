package tumblr

import (
	"net/url"
	"path"
)

type Blog struct {
	Basename string
}

const (
	urlBase = "http://api.tumblr.com/v2/"
)
