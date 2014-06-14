package tumblr

import (
	"testing"
	"fmt"
)

func TestPosts(t *testing.T) {
	b := makeTumblr().NewBlog("lacecard.tumblr.com")
	// Check for unique post
	params := PostRequestParams{
		Id: int64(76803575816),
	}
	posts, err := b.Posts(params)
	if err != nil {
		t.Error(err)
		return
	}
	if len(posts) != 1 {
		t.Errorf("Specified ID, expecting one post, got %d", len(posts))
		return
	}
}
