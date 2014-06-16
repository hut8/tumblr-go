package tumblr

// Defines each subtype of Post (see consts below) and factory methods

import (
	"encoding/json"
	//"fmt"
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
	Posts       []Post // A conjunction of the below
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
	rawPosts := []json.RawMessage{}
	//jsonSource, err := r.MarshalJSON()
	//fmt.Printf("%s\n", string(jsonSource))
	err := json.Unmarshal(*r, &rawPosts)
	if err != nil {
		return nil, err
	}
	pc := &PostCollection{}
	// Append the post to the right field
	for _, rp := range rawPosts {
		// Extract most generic sections first
		var p PostBase
		err = json.Unmarshal(rp, &p)
		if err != nil {
			return nil, err
		}

		// Based on the type of the post, create a TypePost (sp = specific post)
		switch p.PostType() {
		case Text:
			var sp TextPost
			json.Unmarshal(rp, &sp)
			pc.TextPosts = append(pc.TextPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Quote:
			var sp QuotePost
			json.Unmarshal(rp, &sp)
			pc.QuotePosts = append(pc.QuotePosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Link:
			var sp LinkPost
			json.Unmarshal(rp, &sp)
			pc.LinkPosts = append(pc.LinkPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Answer:
			var sp AnswerPost
			json.Unmarshal(rp, &sp)
			pc.AnswerPosts = append(pc.AnswerPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Video:
			var sp VideoPost
			json.Unmarshal(rp, &sp)
			pc.VideoPosts = append(pc.VideoPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Audio:
			var sp AudioPost
			json.Unmarshal(rp, &sp)
			pc.AudioPosts = append(pc.AudioPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Photo:
			var sp PhotoPost
			json.Unmarshal(rp, &sp)
			pc.PhotoPosts = append(pc.PhotoPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
		case Chat:
			var sp ChatPost
			json.Unmarshal(rp, &sp)
			pc.ChatPosts = append(pc.ChatPosts, sp)
			pc.Posts = append(pc.Posts, &sp)
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
	Notes       []NoteData
	TotalPosts  int64 // total posts in result set for pagination
}

// Accessors for the common fields of a Post
type Post interface {
	PostBlogName() string
	PostId() int64
	PostPostURL() string
	PostTimestamp() int64
	PostType() PostType
	PostDate() string
	PostFormat() string
	PostReblogKey() string
	PostTags() []string
	PostBookmarklet() bool
	PostMobile() bool
	PostSourceURL() string
	PostSourceTitle() string
	PostLiked() bool
	PostState() string // published, ueued, draft, private
	PostNotes() []NoteData
	PostTotalPosts() int64 // total posts in result set for pagination
}

func (p *PostBase) PostBlogName() string    { return p.BlogName }
func (p *PostBase) PostId() int64           { return p.Id }
func (p *PostBase) PostPostURL() string     { return p.PostURL }
func (p *PostBase) PostType() PostType      { return TypeOfPost(p.Type) }
func (p *PostBase) PostTimestamp() int64    { return p.Timestamp }
func (p *PostBase) PostDate() string        { return p.Date }
func (p *PostBase) PostFormat() string      { return p.Format }
func (p *PostBase) PostReblogKey() string   { return p.ReblogKey }
func (p *PostBase) PostTags() []string      { return p.Tags }
func (p *PostBase) PostBookmarklet() bool   { return p.Bookmarklet }
func (p *PostBase) PostMobile() bool        { return p.Mobile }
func (p *PostBase) PostSourceURL() string   { return p.SourceURL }
func (p *PostBase) PostSourceTitle() string { return p.SourceTitle }
func (p *PostBase) PostLiked() bool         { return p.Liked }
func (p *PostBase) PostState() string       { return p.State }
func (p *PostBase) PostNotes() []NoteData   { return p.Notes }
func (p *PostBase) PostTotalPosts() int64   { return p.TotalPosts }

// Text post
type TextPost struct {
	PostBase
	Title string
	Body  string
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
	Players []EmbedObjectData
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

// General notes information
// {
// 	  "timestamp": "1401041794", // or an integer! lol
// 	  "blog_name": "nalisification",
// 	  "blog_url": "http://nalisification.tumblr.com/",
// 	  "post_id": "1234",
// 	  "type": "reblog"
// },
type NoteData struct {
	Timestamp interface{} // this is either a string or an integer :(
	BlogName  string      `json:"blog_name"`
	BlogURL   string      `json:"blog_url"`
	PostID    string      `json:"post_id"` // wtf
	Type      string      `json:"type"`    // reblog, like, post, ...?
}
