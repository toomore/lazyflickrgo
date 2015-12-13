package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/toomore/lazyflickrgo/flickr"
	"github.com/toomore/lazyflickrgo/jsonstruct"
)

var (
	userID  = flag.String("userid", "", "User number ID")
	albumID = flag.String("albumid", "", "Album/Set number ID")
	groupID = flag.String("groupid", "", "Group number ID")
	apikey  = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret  = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	shareN  = flag.Int64("n", 6, "Per share num")
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var wg sync.WaitGroup

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
	if num <= *shareN {
		*shareN = num
	}
	wg.Add(int(*shareN))
	for _, val := range r.Perm(int(num))[:*shareN] {
		photo := albumdata.Photoset.Photos.Photo[val]
		go func(photo jsonstruct.Photo, groupID *string, val int) {
			log.Println(val, photo.ID, photo)
			log.Printf("%+v", f.GroupsPoolsAdd(*groupID, photo.ID))
			wg.Done()
		}(photo, groupID, val)
	}
	wg.Wait()
}
