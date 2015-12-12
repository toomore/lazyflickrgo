// Package request support Get/Post request.
package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/toomore/lazyflickrgo/jsonstruct"
)

// Request struct
type Request struct {
	URL  *url.URL
	args *url.Values
}

// NewRequest is to new a request.
func NewRequest(URL string, APIKey string) *Request {
	args := &url.Values{}

	// Default args.
	args.Set("api_key", APIKey)
	args.Set("format", "json")
	args.Set("nojsoncallback", "1")

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(errors.New("URL format fail"))
	}
	return &Request{
		URL:  url,
		args: args,
	}
}

// Get is Get method request.
func (r Request) Get(Args map[string]string) *http.Response {
	for key, val := range Args {
		r.args.Add(key, val)
	}
	r.URL.RawQuery = r.args.Encode()
	log.Println("Get: ", r.URL.String())
	resp, err := http.Get(r.URL.String())
	if err != nil {
		log.Fatalln(err)
	}

	return resp

}

// PhotosSearch is "flickr.photos.search"
//
// https://www.flickr.com/services/api/flickr.photos.search.html
func (r Request) PhotosSearch(Args map[string]string) jsonstruct.PhotosSearch {
	Args["method"] = "flickr.photos.search"

	resp := r.Get(Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)

	var data jsonstruct.PhotosSearch
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}
