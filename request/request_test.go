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

	data := t.PhotosSearch(args)

	for i, vals := range data.Photos.Photo {
		log.Printf("%02d %+v\n", i, vals)
		//log.Printf("https://www.flickr.com/photos/%s/%s\n", vals.Owner, vals.ID)
	}
}
