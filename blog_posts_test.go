package tumblr

import (
	"testing"
)

func TestPosts(t *testing.T) {
	b := makeTumblr().NewBlog("lacecard.tumblr.com")
	// Check for unique post
	params := PostRequestParams{
		Id: int64(76803575816),
		NotesInfo: true,
	}
	pc, err := b.Posts(params)
	if err != nil {
		t.Error(err)
		return
	}
	if len(pc.Posts) != 1 {
		t.Errorf("Specified ID, expecting one post, got %d", len(pc.Posts))
		return
	}
	if len(pc.Posts[0].PostNotes()) == 0 {
		t.Errorf("Did not find notes that I expected to find")
		return
	}
	// TODO Add test to make sure that the notes are actually deserialized properly
}
