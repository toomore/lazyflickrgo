package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/toomore/lazyflickrgo/flickr"
)

var (
	userID  = flag.String("userid", "", "User number ID")
	albumID = flag.String("albumid", "", "Album/Set number ID")
	groupID = flag.String("groupid", "", "Group number ID")
	apikey  = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret  = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	shareN  = flag.Int("n", 6, "Per share num")
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	f := flickr.NewFlickr(*apikey, *secret)
	f.AuthToken = os.Getenv("FLICKRUSERTOKEN")
	albumdata := f.PhotosetsGetPhotos(*albumID, *userID)
	var num int64
	total, _ := strconv.ParseInt(albumdata.Photoset.Photos.Total, 10, 32)
	if albumdata.Photoset.Photos.Perpage <= total {
		num = albumdata.Photoset.Photos.Perpage
	} else {
		num = total
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, val := range r.Perm(int(num))[:*shareN] {
		photo := albumdata.Photoset.Photos.Photo[val]
		log.Println(val, photo.ID, photo)
		log.Printf("%+v", f.GroupsPoolsAdd(*groupID, photo.ID))
	}
}
