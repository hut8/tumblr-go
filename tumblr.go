package tumblr

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TumblrAPIResponse struct {
	Meta     Meta
	Response interface{}
}

type Meta struct {
	Status int64
	Msg    string
}

type APICredentials struct {
	Key string
	Secret string
}

type LimitOffset struct {
	Limit  int
	Offset int
}

type Tumblr struct {
	Credentials APICredentials
}

// API Functions

func callAPI(u *url.URL) (*simplejson.Json, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}

	// Handle Meta
	statCode, err := json.Get("meta").Get("status").Int64()
	if err != nil {
		return nil, err
	}
	statMsg, err := json.Get("meta").Get("msg").String()
	if err != nil {
		return nil, err
	}
	meta := Meta{
		Status: statCode,
		Msg: statMsg,
	}
	if meta.Status != 200 {
		err = fmt.Errorf("tumblr API responded with HTTP status %d: %s",
			meta.Status,
			meta.Msg)
		return nil, err
	}

	res := json.Get("response")

	fmt.Printf("response for %v:\n%# v\n", u, pretty.Formatter(res))

	return res, nil
}

func (t Tumblr) NewBlog(baseHostname string) (Blog) {
	return Blog{
		BaseHostname: baseHostname,
		t: t,
	}
}

type Post interface{}

const (
	apiBaseURL = "http://api.tumblr.com/v2/"
)

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

func (t Tumblr) apiURL() (*url.URL, error) {
	url, err := url.Parse(apiBaseURL)
	if err != nil {
		return nil, err
	}
	addCredentials(url, t.Credentials)
	return url, nil
}

// Request Parameter Types

func addCredentials(url *url.URL, credentials APICredentials) {
	vals := url.Query()
	vals.Set("api_key", credentials.Key)
	url.RawQuery = vals.Encode()
}

func addLimitOffset(url *url.URL, params LimitOffset) {
	// Limit
	if params.Limit != 0 {
		url.Query().Set("limit", string(params.Limit))
	}

	// Offset
	if params.Offset != 0 {
		url.Query().Set("offset", string(params.Offset))
	}
}
