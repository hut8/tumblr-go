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

	likesCount, err := res.Get("liked_count").Int64()
	posts := make([]Post, 0, likesCount)
	rawPosts, err := res.Get("liked_posts").Array()
	for _, _ = range rawPosts {
		posts = append(posts, Post{})
	}

	// Parse out post objects
	likes := &BlogLikes{
	 	TotalCount: likesCount,
		Likes: posts,
	}

	return likes, nil
}

type BlogLikes struct {
	Likes      []Post
	TotalCount int64
}
