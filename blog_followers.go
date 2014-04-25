package tumblr

// Blog followers
func (blog Blog) Followers(params LimitOffset) ([]Blog, error) {
	url, err := blog.blogEntityURL("followers")
	if err != nil {
		return err
	}
	addLimitOffset(url, params)
	// TODO send to API
	// TODO parse result
	return nil, nil
}

type BlogFollowers struct {
	TotalUsers int64
	Users      []FollowingUser
}

type FollowingUser struct {
	User
	Following bool // Is the following reciprocal?
}
