package request

import (
	"log"
	"os"
	"testing"
)

func TestNewRequest(*testing.T) {
	t := NewRequest(
		"https://api.flickr.com/services/rest/",
		os.Getenv("FLICKRAPI"))

	log.Printf("%+v", t)

	args := make(map[string]string)
	args["user_id"] = os.Getenv("FLICKRUSER")
	args["tags"] = "lomo,kodak"
	args["tag_mode"] = "all"
	args["sort"] = "date-posted-desc"

	t.PhotosSearch(args)
}
