package flickr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strconv"
	"sync"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

const perPage = 500

// PhotosetsGetPhotos get photos from set/album.
func (f Flickr) PhotosetsGetPhotos(photosetID string, userID string, page int) jsonstruct.PhotosetsGetPhotos {
	args := make(map[string]string)
	args["method"] = "flickr.photosets.getPhotos"
	args["photoset_id"] = photosetID
	args["user_id"] = userID
	args["per_page"] = strconv.Itoa(perPage)
	args["page"] = strconv.Itoa(page)

	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data jsonstruct.PhotosetsGetPhotos
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// PhotosetsGetPhotosAll get all pages data.
func (f Flickr) PhotosetsGetPhotosAll(photosetID string, userID string) []jsonstruct.PhotosetsGetPhotos {
	photosetInfo := f.PhotosetsGetInfo(photosetID, userID)
	pages := int(math.Ceil(float64(photosetInfo.Photoset.Photos) / perPage))
	result := make([]jsonstruct.PhotosetsGetPhotos, pages)

	var wg sync.WaitGroup
	wg.Add(pages)

	go func() {
		runtime.Gosched()
		for i := 0; i < pages; i++ {
			go func(i int) {
				runtime.Gosched()
				defer wg.Done()
				result[i] = f.PhotosetsGetPhotos(photosetID, userID, i+1)
			}(i)
		}
	}()
	wg.Wait()
	return result
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
