package tumblr

// Defines each subtype of Post (see consts below) and factory methods

// Post Types
const (
	Text   = "text"
	Quote  = "quote"
	Link   = "link"
	Answer = "answer"
	Video  = "video"
	Audio  = "audio"
	Photo  = "photo"
	Chat   = "chat"
)

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

type TextPost struct {
	Post
	Title string
	Body  string
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
	Title       string
	URL         string
	Description string
}
