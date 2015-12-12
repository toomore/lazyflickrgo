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
	Args *url.Values
}

func NewRequest(URL string, Method string, Api_key string) *request {
	args := &url.Values{}

	args.Add("api_key", Api_key)
	args.Add("method", Method)
	args.Add("format", "json")
	args.Add("nojsoncallback", "1")

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(errors.New("URL format fail"))
	}
	return &request{
		URL:  url,
		Args: args,
	}
}

func (r request) Get(Args map[string]string) {
	r.URL.RawQuery = r.Args.Encode()
	log.Println("Get: ", r.URL.String())
	resp, err := http.Get(r.URL.String())
	if err != nil {
		log.Fatalln(err)
	}
	jsonData, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", jsonData)

	var data jsonstruct.PhotosSearch
	json.Unmarshal(jsonData, &data)
	log.Printf("%+v", data)
	for i, vals := range data.Photos.Photo {
		log.Println(i, vals)
		//log.Printf("https://www.flickr.com/photos/%s/%s\n", vals.Owner, vals.ID)
	}
}
