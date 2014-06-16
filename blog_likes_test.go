package tumblr

import (
	"testing"
//	"fmt"
)

func TestLikes(t *testing.T) {
	// Note this is an improper "base hostname", which is usually what we get
	// this API makes very little sense.
	b := makeTumblr().NewBlog("http://lacecard.tumblr.com/")
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
