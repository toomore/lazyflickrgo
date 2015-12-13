// Package flickr for api.
package flickr

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

// Flickr struct
type Flickr struct {
	args map[string]string
}

// NewFlickr is to new a request.
func NewFlickr(APIKey string) *Flickr {
	args := make(map[string]string)

	// Default args.
	args["format"] = "json"
	args["nojsoncallback"] = "1"
	args["api_key"] = APIKey

	return &Flickr{
		args: args,
	}
}

// HTTPGet method request.
func (f Flickr) HTTPGet(URL string, Args map[string]string) *http.Response {
	for key, val := range f.args {
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

// HTTPPost method request.
func (f Flickr) HTTPPost(urlpath string, Data map[string]string) *http.Response {
	for key, val := range f.args {
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
func (f Flickr) PhotosSearch(Args map[string]string) jsonstruct.PhotosSearch {
	Args["method"] = "flickr.photos.search"

	resp := f.HTTPGet(utils.APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.PhotosSearch
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// GroupsPoolsAdd add photo to a groups.
func (f Flickr) GroupsPoolsAdd(GroupsID string, PhotosID string) jsonstruct.Common {
	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = GroupsID
	data["photo_id"] = PhotosID
	data["auth_token"] = os.Getenv("FLICKRUSERTOKEN")

	resp := f.HTTPPost(utils.APIURL, data)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result jsonstruct.Common
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}

// AuthGetFrob to get Frob link.
func (f Flickr) AuthGetFrob() jsonstruct.AuthGetFrob {
	Args := map[string]string{"method": "flickr.auth.getFrob"}
	resp := f.HTTPGet(utils.APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.AuthGetFrob
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// AuthGetToken to get user auth token.
func (f Flickr) AuthGetToken(frob string) jsonstruct.AuthGetToken {
	args := make(map[string]string)
	args["method"] = "flickr.auth.getToken"
	args["frob"] = frob

	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", jsonData)
	defer resp.Body.Close()

	var data jsonstruct.AuthGetToken
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}

	return data
}
