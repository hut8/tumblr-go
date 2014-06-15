package tumblr

// Defines each subtype of Post (see consts below) and factory methods

import (
	"encoding/json"
)

// Post Types
type PostType int

const (
	Unknown PostType = iota
	Text
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
	d := Unknown
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

type PostCollection struct {
	Posts       []Post // A combination of the below
	TextPosts   []TextPost
	QuotePosts  []QuotePost
	LinkPosts   []LinkPost
	AnswerPosts []AnswerPost
	VideoPosts  []VideoPost
	AudioPosts  []AudioPost
	PhotoPosts  []PhotoPost
	ChatPosts   []ChatPost
}

// Constructs a PostCollection of typed Posts given the json.RawMessage
// of "response":"posts" which must be an array
func NewPostCollection(r *json.RawMessage) (*PostCollection, error) {
	posts := []Post{}
	err := json.Unmarshal(*r, posts)
	if err != nil {
		return nil, err
	}
	pc := &PostCollection{}
	// Append the post to the right field
	for _, p := range posts {
		switch p.Type {
		case Text:
		case Quote:
		case Link:
		case Answer:
		case Video:
		case Audio:
		case Photo:
		case Chat:
		}
	}
	return pc, nil
}

// Stuff in the "response":"posts" field
type PostBase struct {
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

// Accessors for the common fields of a Post
type Post interface {
	BlogName() string
	Id() int64
	PostURL() string
	Timestamp() int64
	Type() PostType
	Date() string
	Format() string
	ReblogKey() string
	Tags() []string
	Bookmarklet() bool
	Mobile() bool
	SourceURL() string
	SourceTitle() string
	Liked() bool
	State() string     // published, ueued, draft, private
	TotalPosts() int64 // total posts in result set for pagination
}

func (p *Post) BlogName() string    { return p.BlogName }
func (p *Post) Id() int64           { return p.Id }
func (p *Post) PostURL() string     { return p.PostURL }
func (p *Post) Type() PostType      { return TypeOfPost(p.Type) }
func (p *Post) Timestamp() int64    { return p.Timestamp }
func (p *Post) Date() string        { return p.Date }
func (p *Post) Format() string      { return p.Format }
func (p *Post) ReblogKey() string   { return p.ReblogKey }
func (p *Post) Tags() []string      { return p.Tags }
func (p *Post) Bookmarklet() bool   { return p.Bookmarklet }
func (p *Post) Mobile() bool        { return p.Mobile }
func (p *Post) SourceURL() string   { return p.SourceURL }
func (p *Post) SourceTitle() string { return p.SourceTitle }
func (p *Post) Liked() bool         { return p.Liked }
func (p *Post) State() string       { return p.State }
func (p *Post) TotalPosts() int64   { return p.TotalPosts }

// Text post
type TextPost struct {
	PostBase
	Title string
	Body  string
}

func NewTextPost(r json.RawMessage) (*TextPost, error) {
	p := &TextPost{}
	err := json.Unmarshal(r, &p)
	return p, err
}

// Photo post
type PhotoPost struct {
	PostBase
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
	PostBase
	Text   string
	Source string
}

// Link post
type LinkPost struct {
	PostBase
	Title       string
	URL         string
	Description string
}

// Chat post
type ChatPost struct {
	PostBase
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
	PostBase
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
	PostBase
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
	PostBase
	AskingName string
	AskingURL  string
	Question   string
	Answer     string
}
