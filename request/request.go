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

type request struct {
	URL  *url.URL
	args *url.Values
}

func NewRequest(URL string, Api_key string) *request {
	args := &url.Values{}

	// Default args.
	args.Set("api_key", Api_key)
	args.Set("format", "json")
	args.Set("nojsoncallback", "1")

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(errors.New("URL format fail"))
	}
	return &request{
		URL:  url,
		args: args,
	}
}

func (r request) Get(Args map[string]string) *http.Response {
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

func (r request) PhotosSearch(Args map[string]string) jsonstruct.PhotosSearch {
	Args["method"] = "flickr.photos.search"

	resp := r.Get(Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	//log.Printf("%s", jsonData)

	var data jsonstruct.PhotosSearch
	json.Unmarshal(jsonData, &data)
	//log.Printf("%+v", data)
	for i, vals := range data.Photos.Photo {
		log.Println(i, vals)
		//log.Printf("https://www.flickr.com/photos/%s/%s\n", vals.Owner, vals.ID)
	}
	return data
}
