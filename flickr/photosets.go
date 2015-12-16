package flickr

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// PhotosetsGetPhotos get photos from set/album.
func (f Flickr) PhotosetsGetPhotos(photosetID string, userID string) jsonstruct.PhotosetsGetPhotos {
	args := make(map[string]string)
	args["method"] = "flickr.photosets.getPhotos"
	args["photoset_id"] = photosetID
	args["user_id"] = userID

	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.PhotosetsGetPhotos
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// PhotosetsGetInfo get album / set info.
func (f Flickr) PhotosetsGetInfo(photosetID string, userID string) jsonstruct.PhotosetsGetInfo {
	args := make(map[string]string)
	args["method"] = "flickr.photosets.getInfo"
	args["photoset_id"] = photosetID
	args["user_id"] = userID

	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.PhotosetsGetInfo
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}
