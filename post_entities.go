package tumblr

// Defines each subtype of Post (see consts below) and factory methods

import (
	"encoding/json"
)

// Post Types
type PostType int

const (
	Text = iota
	Quote
	Link
	Answer
	Video
	Audio
	Photo
	Chat
)

// Return the PostType of the type described in the JSON
func TypeOfPost(t string) PostType {
	var d PostType
	switch t {
	case "text":
		d = Text
	case "quote":
		d = Quote
	case "link":
		d = Link
	case "answer":
		d = Answer
	case "video":
		d = Video
	case "audio":
		d = Audio
	case "photo":
		d = Photo
	case "chat":
		d = Chat
	}
	return d
}

type PostEntity interface {
	Type() PostType
}

// Stuff in the "response":"posts" field
type Post struct {
	BlogName    string
	Id          int64
	PostURL     string
	Type        string
	Timestamp   int64
	Date        string
	Format      string
	ReblogKey   string
	Tags        []string
	Bookmarklet bool
	Mobile      bool
	SourceURL   string
	SourceTitle string
	Liked       bool
	State       string // published, ueued, draft, private
	TotalPosts  int64  // total posts in result set for pagination
}

// Text post
type TextPost struct {
	Post
	Title string
	Body  string
}

func NewTextPost(r json.RawMessage) (*TextPost, error) {
	p := &TextPost{}
	err := json.Unmarshal(r, &p)
	return p, err
}

func (p *TextPost) Type() PostType {
	return Text
}

// Photo post
type PhotoPost struct {
	Post
	Photos  []PhotoData
	Caption string
	Width   int64
	Height  int64
}

// One photo in a PhotoPost
type PhotoData struct {
	Caption  string // photosets only
	AltSizes []AltSizeData
}

// One alternate size of a Photo
type AltSizeData struct {
	Width  int
	Height int
	URL    string
}

// Quote post
type QuotePost struct {
	Post
	Text   string
	Source string
}

// Link post
type LinkPost struct {
	Post
	Title       string
	URL         string
	Description string
}

// Chat post
type ChatPost struct {
	Post
	Title    string
	Body     string
	Dialogue []DialogueData
}

// One component of a conversation in a Dialogue in a Chat
type DialogueData struct {
	Name   string
	Label  string
	Phrase string
}

// Audio post
type AudioPost struct {
	Post
	Caption     string
	Player      string
	Plays       int64
	AlbumArt    string
	Artist      string
	Album       string
	TrackName   string
	TrackNumber int64
	Year        int
}

// Video post - TODO Handle all the different sources - not documented :(
type VideoPost struct {
	Post
	Caption string
	Player  []EmbedObjectData
}

// One embedded video player in a VideoPost
type EmbedObjectData struct {
	Width     int
	EmbedCode string
}

// Answer post
type AnswerPost struct {
	Post
	AskingName string
	AskingURL  string
	Question   string
	Answer     string
}
