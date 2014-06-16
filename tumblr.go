package tumblr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type TumblrAPIResponse struct {
	Meta     Meta
	Response *json.RawMessage
}

type Meta struct {
	Status int64
	Msg    string
}

type APICredentials struct {
	Key    string
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

func callAPI(u *url.URL) (*json.RawMessage, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &TumblrAPIResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	if res.Meta.Status != 200 {
		err = fmt.Errorf("tumblr API responded with HTTP status %d: %s",
			res.Meta.Status,
			res.Meta.Msg)
		return nil, err
	}

	//subJson, _ := res.Response.MarshalJSON()
	//fmt.Printf("response for %v:\n%# v\n%s\n", u, pretty.Formatter(res), string(subJson))

	return res.Response, nil
}

const (
	apiBaseURL = "http://api.tumblr.com/v2/"
)

func (t Tumblr) apiURL() (*url.URL, error) {
	url, err := url.Parse(apiBaseURL)
	if err != nil {
		return nil, err
	}
	addCredentials(url, t.Credentials)
	return url, nil
}

func addCredentials(url *url.URL, credentials APICredentials) {
	vals := url.Query()
	vals.Set("api_key", credentials.Key)
	url.RawQuery = vals.Encode()
}

func addLimitOffset(url *url.URL, params *LimitOffset) {
	// Limit
	vals := url.Query()
	if params.Limit != 0 {
		vals.Set("limit", strconv.Itoa(params.Limit))
	}

	// Offset
	if params.Offset != 0 {
		vals.Set("offset", strconv.Itoa(params.Offset))
	}
	url.RawQuery = vals.Encode()
}
