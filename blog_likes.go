package tumblr

// Posts liked by a blog
func (blog Blog) Likes(params LimitOffset) (*BlogLikes, error) {
	url, err := blog.blogEntityURL("likes")
	if err != nil {
		return nil, err
	}
	addLimitOffset(url, params)

	return nil, nil
}

type BlogLikes struct {
	Likes      []Post
	TotalCount int64
}
