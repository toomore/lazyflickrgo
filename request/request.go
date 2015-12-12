package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", data)
}
