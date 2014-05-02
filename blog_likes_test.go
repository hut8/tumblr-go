package tumblr

import (
	"testing"
	"fmt"
	"os"
)

func TestLikes(t *testing.T) {
	b := Blog{
		BaseHostname: "lacecard",
		Credentials: APICredentials{
			Key: os.Getenv("MUMBLR_API_KEY"),
		},
	}
	l, err := b.Likes(LimitOffset{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("total count: %d", l.TotalCount)
}
