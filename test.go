package tumblr

import (
	"os"
)

func makeTumblr() Tumblr {
	key := os.Getenv("TUMBLR_API_KEY")
	if key == "" {
		panic("Set your TUMBLR_API_KEY environment variable plz")
	}
	return Tumblr{
		Credentials: APICredentials{
			Key: key,
		},
	}
}
