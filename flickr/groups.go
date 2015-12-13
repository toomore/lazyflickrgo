package flickr

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// GroupsPoolsAdd add photo to a groups.
func (f Flickr) GroupsPoolsAdd(GroupsID string, PhotosID string) jsonstruct.Common {
	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = GroupsID
	data["photo_id"] = PhotosID
	data["auth_token"] = f.secretKey

	resp := f.HTTPPost(utils.APIURL, data)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result jsonstruct.Common
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}
