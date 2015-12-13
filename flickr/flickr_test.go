package flickr

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/toomore/lazyflickrgo/utils"
)

func getFlickr() *Flickr {
	t := NewFlickr(os.Getenv("FLICKRAPIKEY"))

	log.Printf("%+v\n", t)

	return t
}

func TestFlickr_PhotosSearch(*testing.T) {
	t := getFlickr()

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

	log.Printf("%+v", data)
}

func TestFlickr_Post(*testing.T) {
	t := getFlickr()

	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = os.Getenv("FLICKRGROUPID")
	data["photo_id"] = os.Getenv("FLICKRPHOTOID")
	data["auth_token"] = os.Getenv("FLICKRUSERTOKEN")

	resp := t.HttpPost(utils.APIURL, data)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s, %s\n", body, err)
}

func TestFlickr_AuthGetFrob(*testing.T) {
	t := getFlickr()
	getFrob := t.AuthGetFrob()
	log.Printf("%+v", getFrob)
	log.Println(getFrob.GetTokenURL())
}

func TestFlickr_GetToken(*testing.T) {
	t := getFlickr()
	log.Printf("%+v", t.AuthGetToken("72157660016985653-8e43466dd79cd0b2-812975"))
}

func TestFlickr_GroupsPoolsAdd(*testing.T) {
	t := getFlickr()
	log.Printf("%+v\n",
		t.GroupsPoolsAdd(os.Getenv("FLICKRGROUPID"), "21111643239"),
	)
}
