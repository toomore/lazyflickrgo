package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Pinterest type struct
type Pinterest struct {
	AccessToken string
}

// APIURL pinterest base.
const APIURL = "https://api.pinterest.com"

// Get HTTP GET
func (p Pinterest) Get(path string, params url.Values) (*http.Response, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("access_token", p.AccessToken)
	url := fmt.Sprintf("%s%s?%s", APIURL, path, params.Encode())
	log.Printf("Get %s\n", url)
	return http.Get(url)
}

// Me Get
func (p Pinterest) Me() {
	resp, err := p.Get("/v1/me/", nil)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s\n", body)
	}
}

func main() {
	pin := &Pinterest{AccessToken: os.Getenv("PINTEREST_TOKEN")}
	pin.Me()
}
