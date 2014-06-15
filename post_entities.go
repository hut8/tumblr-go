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

type PhotoPost struct {
	Post
	Photos []PhotoData
}

type PhotoData struct {
}
