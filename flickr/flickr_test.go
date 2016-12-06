package flickr

import (
	"log"
	"os"
	"testing"

	"github.com/toomore/lazyflickrgo/utils"
)

func getFlickr() *Flickr {
	t := NewFlickr(os.Getenv("FLICKRAPIKEY"), os.Getenv("FLICKRSECRET"))

	log.Printf("Flickr info: %+v\n", t)

	return t
}

func TestFlickr_PhotosSearch(*testing.T) {
	t := getFlickr()

	args := make(map[string]string)
	args["user_id"] = os.Getenv("FLICKRUSER")
	//args["tags"] = "lomo,kodak"
	args["tags"] = "nikon"
	args["tag_mode"] = "all"
	args["sort"] = "date-posted-desc"

	datapages := t.PhotosSearch(args)

	for _, data := range datapages {
		log.Printf(">>>>>>>>>>>>> %d \n", data.Photos.Page)
		log.Printf(">>>>>>>>>>>>>Perpage %d \n", len(data.Photos.Photo))
		//for i, vals := range data.Photos.Photo {
		//	log.Printf("%02d %+v\n", i, vals)
		//	//log.Printf("https://www.flickr.com/photos/%s/%s\n", vals.Owner, vals.ID)
		//}
		//log.Printf("%+v", data)
	}
	log.Println(args)

}

func TestFlickr_Post(*testing.T) {
	t := getFlickr()

	data := make(map[string]string)
	data["method"] = "flickr.groups.pools.add"
	data["group_id"] = os.Getenv("FLICKRGROUPID")
	data["photo_id"] = os.Getenv("FLICKRPHOTOID")
	data["auth_token"] = os.Getenv("FLICKRUSERTOKEN")

	resp := t.HTTPPost(utils.APIURL, data)
	log.Printf("%s\n ", resp)
}

func TestFlickr_AuthGetFrob(*testing.T) {
	t := getFlickr()
	getFrob := t.AuthGetFrob()
	log.Printf("%+v", getFrob)
	log.Println(getFrob.GetTokenURL(os.Getenv("FLICKRAPIKEY"), os.Getenv("FLICKRSECRET")))
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

func TestFlickr_PhotosetsGetPhotos(*testing.T) {
	t := getFlickr()
	data := t.PhotosetsGetPhotos("72157656102091084", os.Getenv("FLICKRUSER"), 1)
	log.Printf("%+v\n", data)
	log.Println(len(data.Photoset.Photo))
}

func TestFlickr_PhotosetsGetPhotosAll(*testing.T) {
	t := getFlickr()
	data := t.PhotosetsGetPhotosAll("72157656102091084", os.Getenv("FLICKRUSER"))
	log.Printf("%+v\n", data)
	for _, val := range data {
		log.Println(val.Photoset.Photos.Page)
		log.Println(len(val.Photoset.Photos.Photo))
	}
}

func TestFlickr_GroupsGetInfo(*testing.T) {
	t := getFlickr()
	t.AuthToken = os.Getenv("FLICKRUSERTOKEN")
	log.Printf("%+v\n",
		t.GroupsGetInfo("", "japan_directory_nihon", false).Group.Throttle,
	)
	log.Printf("%+v\n",
		t.GroupsGetInfo("11526962@N00", "", false).Group.Throttle,
	)
	log.Printf("%+v\n",
		t.GroupsGetInfo("14431758@N00", "", true).Group.Throttle,
	)
}

func TestFlickr_PhotosetsGetInfo(*testing.T) {
	t := getFlickr()
	log.Printf("%+v\n",
		t.PhotosetsGetInfo("72157656102091084", os.Getenv("FLICKRUSER")),
	)
}

func TestFlickr_PhotosGetInfo(*testing.T) {
	t := getFlickr()
	log.Printf("%+v\n",
		t.PhotosGetInfo("23544438000"),
	)
}

func TestFlickr_PhotosLicensesGetInfo(*testing.T) {
	t := getFlickr()
	log.Printf("%+v\n",
		t.PhotosLicensesGetInfo(),
	)
}

func TestFlickr_PeopleFindByEmail(*testing.T) {
	t := getFlickr()
	log.Printf("%+v\n",
		t.PeopleFindByEmail(""),
	)
	log.Printf("%+v\n",
		t.PeopleFindByEmail("toomore0929@gmail.com"),
	)
	log.Printf("%+v\n",
		t.PeopleFindByUsername("toomore"),
	)
}

func TestFlickr_PeopleGetGroups(*testing.T) {
	t := getFlickr()
	t.AuthToken = os.Getenv("FLICKRUSERTOKEN")
	log.Printf("%+v\n",
		t.PeopleGetGroups("92438116@N00", ""),
	)
	//for num, val := range t.PeopleGetGroups("92438116@N00", "").Groups.Group {
	//	log.Printf("%d %s \n", num, val.Iconserver)
	//}
}
