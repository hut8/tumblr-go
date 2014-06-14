package tumblr

import (
	"net/url"
	"path"
)

func validateAvatarSize(size int32) bool {
	switch size {
	case 0, 16, 24, 30, 40, 48, 64, 96, 128, 512:
		return true
	default:
		return false
	}
}

func (blog *Blog) Avatar(size int32) (*url.URL, error) {
	url, err := blog.blogEntityURL("avatar")
	if err != nil {
		return nil, err
	}

	if size != 0 {
		if !validateAvatarSize(size) {
			return nil, TumblrError{"Invalid Avatar size specified"}
		}
		url.Path = path.Join(url.Path, string(size))
	}

	// TODO send to API
	// TODO parse results in "avatar_url"
	return nil, nil
}
