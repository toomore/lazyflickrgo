// Package request support Get/Post request.
package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

const (
	// APIURL Flickr API
	APIURL = "https://api.flickr.com/services/rest/"
	// AUTHURL Flickr Auth URL
	AUTHURL = "http://flickr.com/services/auth/"
)

// Request struct
type Request struct {
	args map[string]string
}

// NewRequest is to new a request.
func NewRequest(APIKey string) *Request {
	args := make(map[string]string)

	// Default args.
	args["format"] = "json"
	args["nojsoncallback"] = "1"
	args["api_key"] = APIKey

	return &Request{
		args: args,
	}
}

// Get method request.
func (r Request) Get(URL string, Args map[string]string) *http.Response {
	for key, val := range r.args {
		Args[key] = val
	}

	r.args["api_sig"] = utils.Sign(Args)

	query := url.Values{}
	for key, val := range Args {
		query.Set(key, val)
	}

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(err)
	}
	url.RawQuery = query.Encode()
	log.Println("Get: ", url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

// Post method request.
func (r Request) Post(urlpath string, Data map[string]string) *http.Response {
	for key, val := range r.args {
		Data[key] = val
	}
	log.Printf("Post: %+v %s", r.args, urlpath)

	query := url.Values{}
	for key, val := range Data {
		query.Set(key, val)
	}

	resp, err := http.PostForm(urlpath, query)
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

	resp := r.Get(APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)

	var data jsonstruct.PhotosSearch
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

//func (r Request) AuthGetFrob() string {
//	r.
