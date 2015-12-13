// Package request support Get/Post request.
package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
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

	Args["api_sig"] = utils.Sign(Args)

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

	Data["api_sig"] = utils.Sign(Data)

	log.Printf("Post: %+v %s", Data, urlpath)

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

// PhotosSearch search photos.
//
// https://www.flickr.com/services/api/flickr.photos.search.html
func (r Request) PhotosSearch(Args map[string]string) jsonstruct.PhotosSearch {
	Args["method"] = "flickr.photos.search"

	resp := r.Get(utils.APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.PhotosSearch
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// GroupsPoolsAdd add photo to a groups.
func (r Request) GroupsPoolsAdd(GroupsID string, PhotosID string) jsonstruct.Common {
	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = GroupsID
	data["photo_id"] = PhotosID
	data["auth_token"] = os.Getenv("FLICKRUSERTOKEN")

	resp := r.Post(utils.APIURL, data)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result jsonstruct.Common
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}

// AuthGetFrob to get Frob link.
func (r Request) AuthGetFrob() jsonstruct.AuthGetFrob {
	Args := map[string]string{"method": "flickr.auth.getFrob"}
	resp := r.Get(utils.APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.AuthGetFrob
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// GetToken to get user auth token.
func (r Request) GetToken(frob string) jsonstruct.AuthGetToken {
	args := make(map[string]string)
	args["method"] = "flickr.auth.getToken"
	args["frob"] = frob

	resp := r.Get(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", jsonData)
	defer resp.Body.Close()

	var data jsonstruct.AuthGetToken
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}

	return data
}
