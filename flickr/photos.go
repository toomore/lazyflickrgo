package flickr

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

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
