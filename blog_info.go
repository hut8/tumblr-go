package tumblr

// Blog Info
func (blog Blog) Info() (*BlogInfo, error) {
	url, err := blog.blogEntityURL("info")
	if err != nil {
		return nil, err
	}
	callAPI(*url)
	//res, err := callAPI(url)
	return nil, nil
}

// Type returned by blog.Info()
type BlogInfo struct {
	Title       string
	Posts       int64
	Name        string
	Updated     int64
	Description string
	Ask         bool
	AskAnon     bool
	Likes       int64
}
