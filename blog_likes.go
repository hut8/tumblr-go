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

	res, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	likesCount := int64((res.(map[string]interface{}))["liked_count"].(float64))
	likes := &BlogLikes{
	 	TotalCount: likesCount,
	}

	return likes, nil
}

type BlogLikes struct {
	Likes      []Post
	TotalCount int64
}
