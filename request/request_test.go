package request

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/toomore/lazyflickrgo/utils"
)

func getRequest() *Request {
	t := NewRequest(os.Getenv("FLICKRAPI"))

	log.Printf("%+v\n", t)

	return t
}

func TestRequest_PhotosSearch(*testing.T) {
	t := getRequest()

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

func TestRequest_Post(*testing.T) {
	t := getRequest()

	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = os.Getenv("FLICKRGROUPID")
	data["photo_id"] = os.Getenv("FLICKRPHOTOID")
	data["auth_token"] = os.Getenv("FLICKRUSERTOKEN")

	resp := t.Post(utils.APIURL, data)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s, %s\n", body, err)
}

func TestRequest_AuthGetFrob(*testing.T) {
	t := getRequest()
	getFrob := t.AuthGetFrob()
	log.Printf("%+v", getFrob)
	log.Println(getFrob.GetTokenURL())
}

func TestRequest_GetToken(*testing.T) {
	t := getRequest()
	log.Printf("%+v", t.AuthGetToken("72157660016985653-8e43466dd79cd0b2-812975"))
}

func TestRequest_GroupsPoolsAdd(*testing.T) {
	t := getRequest()
	log.Printf("%+v\n",
		t.GroupsPoolsAdd(os.Getenv("FLICKRGROUPID"), "21111643239"),
	)
}
