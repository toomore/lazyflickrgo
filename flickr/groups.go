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
	data["auth_token"] = f.AuthToken

	resp := f.HTTPPost(utils.APIURL, data)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result jsonstruct.Common
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}

// GroupsGetInfo for search group by id or path.
func (f Flickr) GroupsGetInfo(GroupID string, PathAlias string) jsonstruct.GroupsGetInfo {
	args := make(map[string]string)
	args["method"] = "flickr.groups.getInfo"
	if GroupID != "" {
		args["group_id"] = GroupID
	}

	if PathAlias != "" {
		args["group_path_alias"] = PathAlias
	}
	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result jsonstruct.GroupsGetInfo
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}
