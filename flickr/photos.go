package flickr

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

func readPhotosSerch(f Flickr, args map[string]string, wg *sync.WaitGroup) jsonstruct.PhotosSearch {
	defer wg.Done()
	jsonData := f.HTTPGet(utils.APIURL, args)

	var data jsonstruct.PhotosSearch
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// PhotosSearch search photos.
//
// https://www.flickr.com/services/api/flickr.photos.search.html
func (f Flickr) PhotosSearch(Args map[string]string) []jsonstruct.PhotosSearch {
	Args["method"] = "flickr.photos.search"
	Args["per_page"] = "500"

	var wg sync.WaitGroup
	wg.Add(1)
	data := readPhotosSerch(f, Args, &wg)

	wg.Wait()
	if data.Photos.Pages > 1 {
		result := make([]jsonstruct.PhotosSearch, data.Photos.Pages)
		result[0] = data

		wg.Add(data.Photos.Pages - 1)
		go func() {
			for i := 2; i <= data.Photos.Pages; i++ {
				go func(i int, Args map[string]string) {
					args := make(map[string]string)
					for k, v := range Args {
						args[k] = v
					}
					args["page"] = strconv.Itoa(i)
					result[i-1] = readPhotosSerch(f, args, &wg)
				}(i, Args)
			}
		}()
		wg.Wait()
		return result
	}
	return []jsonstruct.PhotosSearch{data}
}

// PhotosGetInfo get photo info.
//
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
func (f Flickr) PhotosGetInfo(photoID string) jsonstruct.PhotosGetInfo {
	Args := make(map[string]string)
	Args["method"] = "flickr.photos.getInfo"
	Args["photo_id"] = photoID

	jsonData := f.HTTPGet(utils.APIURL, Args)

	var data jsonstruct.PhotosGetInfo
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// PhotosLicensesGetInfo get photo licenses list
//
// https://www.flickr.com/services/api/flickr.photos.licenses.getInfo.html
func (f Flickr) PhotosLicensesGetInfo() jsonstruct.PhotosLicenses {
	Args := make(map[string]string)
	Args["method"] = "flickr.photos.licenses.getInfo"

	jsonData := f.HTTPGet(utils.APIURL, Args)
	var data jsonstruct.PhotosLicenses
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// PhotosGetSizes get photo sizes
//
// https://www.flickr.com/services/api/flickr.photos.getSizes.html
func (f Flickr) PhotosGetSizes(photoID string) jsonstruct.PhotoSizes {
	Args := make(map[string]string)
	Args["method"] = "flickr.photos.getSizes"
	Args["photo_id"] = photoID

	jsonData := f.HTTPGet(utils.APIURL, Args)
	var data jsonstruct.PhotoSizes
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}
