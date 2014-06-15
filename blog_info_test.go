package tumblr

import (
	"testing"
)

func TestInfo(t *testing.T) {
	c := makeTumblr()
	b := c.NewBlog("lacecard.tumblr.com")
	i, err := b.Info()
	if err != nil {
		t.Error(err)
		return
	}
	if i.Updated < 1 {
		t.Errorf("Blog marked as never updated (updated=%d)", i.Updated)
		return
	}
	if i.Likes < 1 {
		t.Errorf("Blog marked as having no likes (likes=%d)", i.Updated)
		return
	}
	if i.Name != "lacecard" {
		t.Errorf("Blog information gives wrong name (name=%s)", i.Name)
		return
	}
}
