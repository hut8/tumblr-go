package tumblr

import (
	"encoding/json"
)

// Posts liked by a blog
func (blog *Blog) Likes(params LimitOffset) (*BlogLikes, error) {
	url, err := blog.blogEntityURL("likes")
	if err != nil {
		return nil, err
	}
	addLimitOffset(url, &params)

	res, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// Decode the response partially
	dr := blogLikesResponse{}
	err = json.Unmarshal(*res, &dr)
	if err != nil {
		return nil, err
	}

	// Make a typed post collection
	pc, err := NewPostCollection(dr.Liked_Posts)
	if err != nil {
		return nil, err
	}

	// Parse out post objects
	likes := &BlogLikes{
	 	TotalCount: dr.Liked_Count,
		Likes: pc,
	}

	return likes, nil
}

type blogLikesResponse struct {
	Liked_Posts *json.RawMessage
	Liked_Count int64
}

type BlogLikes struct {
	Likes      *PostCollection
	TotalCount int64
}
