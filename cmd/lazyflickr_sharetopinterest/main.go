// cmd/lazyflickr_sharetopinterest for share Flickr to Pinterest.
/*
Install:

	go install github.com/toomore/lazyflickrgo/cmd/lazyflickr_sharetopinterest

Usage:

	lazyflickr_sharetopinterest [flags] <flickr photo nsid>[ <flickr photo nsid> ...]

The flags are:
	-board
		Pin board, <username>/<board_name>

	-dryrun
		Show result without post to groups

Required env:

	PINTEREST_TOKEN, FLICKRAPIKEY, FLICKRSECRET

*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/toomore/lazyflickrgo/flickr"
	"github.com/toomore/lazyflickrgo/jsonstruct"
)

// Pinterest type struct
type Pinterest struct {
	AccessToken string
}

// APIURL pinterest base.
const APIURL = "https://api.pinterest.com"

// https://farm{farm-id}.staticflickr.com/{server-id}/{id}_{o-secret}_o.(jpg|gif|png)
const imageFormat = "https://farm%d.staticflickr.com/%s/%s_%s_o.%s"

var (
	board  = flag.String("board", "", "Pin board, <username>/<board_name>")
	dryRun = flag.Bool("dry_run", false, "Dry run")
)

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

// Post HTTP POST
func (p Pinterest) Post(path string, data url.Values) (*http.Response, error) {
	if data == nil {
		data = url.Values{}
	}
	data.Set("access_token", p.AccessToken)
	url := fmt.Sprintf("%s%s", APIURL, path)
	log.Printf("POST %s\n", url)
	return http.PostForm(url, data)
}

// Me Get
func (p Pinterest) Me() {
	resp, err := p.Get("/v1/me/", nil)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s\n", body)
		showRatelimit(resp.Header)
	}
}

// PinsPost for post pins
func (p Pinterest) PinsPost(board, note, link, imageURL string) {
	data := url.Values{}
	data.Set("board", board)
	data.Set("note", note)
	data.Set("link", link)
	data.Set("image_url", imageURL)

	resp, err := p.Post("/v1/pins/", data)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s\n", body)
		showRatelimit(resp.Header)
	}
}

func showRatelimit(Header http.Header) {
	log.Printf("Remaining: %s, Limit: %s", Header.Get("X-Ratelimit-Remaining"), Header.Get("X-Ratelimit-Limit"))
}

func main() {
	flag.Parse()
	pin := &Pinterest{AccessToken: os.Getenv("PINTEREST_TOKEN")}
	//pin.Me()
	f := flickr.NewFlickr(os.Getenv("FLICKRAPIKEY"), os.Getenv("FLICKRSECRET"))

	var (
		imageURL string
		link     string
		note     string
		photo    jsonstruct.PhotosGetInfo
	)
	for _, photoID := range flag.Args() {
		photo = f.PhotosGetInfo(photoID)
		//log.Printf("%+v", photo)

		// Tags
		tags := make([]string, len(photo.Photo.Tags.Tag))
		for i, tag := range photo.Photo.Tags.Tag {
			tags[i] = fmt.Sprintf("#%s", strings.Replace(tag.Raw, " ", "_", -1))
		}
		//log.Println(tags)

		note = fmt.Sprintf("%s - %s - %s",
			photo.Photo.Title.Content,
			strings.Replace(photo.Photo.Description.Content, "\n", " ", -1),
			strings.Join(tags, " "),
		)

		link = photo.Photo.Urls.URL[0].Content
		imageURL = fmt.Sprintf(imageFormat,
			photo.Photo.Farm,
			photo.Photo.Server,
			photo.Photo.ID,
			photo.Photo.Orgsecret,
			photo.Photo.Orgformat)

		log.Println(*board, note, link, imageURL)
		if !*dryRun {
			pin.PinsPost(*board, note, link, imageURL)
		} else {
			log.Println("Dry Run!")
		}
	}
}
