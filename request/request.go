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
	"github.com/toomore/lazyflickrgo/utils"
)

// Request struct
type Request struct {
	URL  *url.URL
	args map[string]string
}

// NewRequest is to new a request.
func NewRequest(URL string, APIKey string) *Request {
	args := make(map[string]string)

	// Default args.
	args["format"] = "json"
	args["nojsoncallback"] = "1"
	args["api_key"] = APIKey

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(errors.New("URL format fail"))
	}
	return &Request{
		URL:  url,
		args: args,
	}
}

// Get method request.
func (r Request) Get(Args map[string]string) *http.Response {
	for key, val := range r.args {
		Args[key] = val
	}

	r.args["api_sig"] = utils.Sign(Args)

	query := url.Values{}
	for key, val := range Args {
		query.Set(key, val)
	}

	r.URL.RawQuery = query.Encode()
	log.Println("Get: ", r.URL.String())
	resp, err := http.Get(r.URL.String())
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

// Post method request.
func (r Request) Post(Data map[string]string) *http.Response {
	for key, val := range r.args {
		Data[key] = val
	}
	log.Printf("Post: %+v %s", r.args, r.URL.String())

	query := url.Values{}
	for key, val := range Data {
		query.Set(key, val)
	}

	resp, err := http.PostForm(r.URL.String(), query)
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
