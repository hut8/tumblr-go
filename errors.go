package tumblr

// Errors
type TumblrError struct {
	Message string
}

func (err TumblrError) Error() string {
	return err.Message
}
