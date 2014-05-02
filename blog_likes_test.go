package tumblr

import (
	"testing"
	"fmt"
	"os"
)

func makeTumblr() Tumblr {
	return Tumblr{
		Credentials: APICredentials{
			Key: os.Getenv("MUMBLR_API_KEY"),
		},
	}
}

func TestLikes(t *testing.T) {
	b := makeTumblr().NewBlog("lacecard.tumblr.com")
	l, err := b.Likes(LimitOffset{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("total count: %d", l.TotalCount)
}
