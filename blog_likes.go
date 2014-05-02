package tumblr

import (
//	"encoding/json"
)

// Posts liked by a blog
func (blog Blog) Likes(params LimitOffset) (*BlogLikes, error) {
	url, err := blog.blogEntityURL("likes")
	if err != nil {
		return nil, err
	}
	addLimitOffset(url, params)

	_, err = callAPI(url)
	if err != nil {
		return nil, err
	}
	// likes := &BlogLikes{
	// 	TotalCount: res["liked_count"].(int64),
	// }

	return nil, nil
}

type BlogLikes struct {
	Likes      []Post
	TotalCount int64
}
