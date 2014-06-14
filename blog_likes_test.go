package tumblr

import (
	"testing"
//	"fmt"
)


func TestLikes(t *testing.T) {
	b := makeTumblr().NewBlog("lacecard.tumblr.com")
	l, err := b.Likes(LimitOffset{})
	if err != nil {
		t.Error(err)
		return
	}
	if l.TotalCount < 1 {
		t.Errorf("There should be some favorites, but there aren't.")
		return
	}
}
