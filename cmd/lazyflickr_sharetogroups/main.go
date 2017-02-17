// cmd/lazyflickr_sharetogroups share photo to groups.
/*
Install:

	go install github.com/toomore/lazyflickrgo/cmd/lazyflickr_sharetogroups

Usage:

	lazyflickr_sharetogroups [flags]

The flags are:

	-apikey
		Flickr API key, default get from env `FLICKRAPIKEY`

	-secret
		Flickr secret, default get from env `FLICKRSECRET`

	-userid
		Flickr userid(nsid), default get from env `FLICKRUSER`

	-albumid
		Album/Set number ID

	-groupid
		Group number ID

	-n
		share photos num. default is 6

	-tags
		Search tags, ',' for split more

	-dryrun
		Show result without post to groups

Example:

share tag:`lomo,japan` to lomo, lomo.tw groups

	lazyflickr_sharetogroups -tags lomo,japan -groupid 40732537997@N01,72262428@N00

*/
package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/toomore/lazyflickrgo/flickr"
	"github.com/toomore/lazyflickrgo/jsonstruct"
)

var (
	userID      = flag.String("userid", os.Getenv("FLICKRUSER"), "User number ID")
	albumID     = flag.String("albumid", "", "Album/Set number ID")
	groupID     = flag.String("groupid", "", "Group number ID")
	photoID     = flag.String("photoid", "", "Photo number ID")
	apikey      = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret      = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	shareN      = flag.Int("n", 6, "Per share num")
	concurrency = flag.Int("c", 20, "send concurrency")
	tags        = flag.String("tags", "", "Search tags, ',' for split more")
	dryrun      = flag.Bool("dryrun", false, "Show result without post to groups")
	info        = color.New(color.Bold, color.FgGreen).SprintfFunc()
	warn        = color.New(color.Bold, color.FgRed).SprintfFunc()
	debugc      = color.New(color.Bold, color.FgHiYellow).SprintfFunc()
	wg          sync.WaitGroup
	photos      []jsonstruct.Photo
	f           *flickr.Flickr
)

func fromSets() []jsonstruct.Photo {
	var result []jsonstruct.Photo
	for _, albumid := range strings.Split(*albumID, ",") {
		if albumid != "" {
			for _, albumdata := range f.PhotosetsGetPhotosAll(albumid, *userID) {
				result = append(result, albumdata.Photoset.Photos.Photo...)
			}
		}
	}

	return result
}

func fromSearch() []jsonstruct.Photo {
	args := make(map[string]string)
	args["tags"] = *tags
	args["tag_mode"] = "all"
	args["sort"] = "date-posted-desc"
	args["per_page"] = "500"
	args["user_id"] = *userID

	searchResult := f.PhotosSearch(args)

	var result []jsonstruct.Photo
	for _, val := range searchResult {
		result = append(result, val.Photos.Photo...)
	}

	return result
}

func addToPool(photo jsonstruct.Photo, groupid string, val int) {
	runtime.Gosched()
	defer wg.Done()
	if *dryrun == false {
		resp := f.GroupsPoolsAdd(groupid, photo.ID)
		if resp.Stat == "ok" {
			log.Println(info("[%s] %s %s", groupid, photo.ID, photo.Title))
		} else {
			log.Println(warn("[%s] %s(%d) %s %s", groupid, resp.Message, resp.Code, photo.ID, photo.Title))
		}
	} else {
		log.Println(debugc("[DryRun] [%s] %s %s", groupid, photo.ID, photo.Title))
	}
}

func addToPoolByPhotoID(groupid string, photoid string, limit chan struct{}) {
	runtime.Gosched()
	defer wg.Done()
	limit <- struct{}{}
	if *dryrun == false {
		resp := f.GroupsPoolsAdd(groupid, photoid)
		if resp.Stat == "ok" {
			log.Println(info("[%s] %s", groupid, photoid))
		} else {
			log.Println(warn("[%s] %s(%d) %s", groupid, resp.Message, resp.Code, photoid))
		}
	} else {
		log.Println(debugc("[DryRun] [%s] %s", groupid, photoid))
	}
	<-limit
}

func send(groupid string, photos []jsonstruct.Photo, randlist []int) {
	runtime.Gosched()
	for _, val := range randlist {
		photo := photos[val]
		log.Println(info("Pick up photo: %d [%s] %+v", val, photo.ID, photo))
		go addToPool(photo, groupid, val)
	}
}

func doShareToGroup() {
	num := len(photos)

	if num == 0 {
		return
	}

	if num <= *shareN {
		*shareN = num
	}

	for _, groupid := range strings.Split(*groupID, ",") {
		wg.Add(*shareN)

		var randlist []int
		startInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
		if (num - *shareN) > 0 {
			startInt = startInt % (num - *shareN)
			randlist = rand.New(rand.NewSource(time.Now().UnixNano())).Perm(num)[startInt : startInt+*shareN]
		} else {
			randlist = rand.New(rand.NewSource(time.Now().UnixNano())).Perm(num)[:*shareN]
		}

		go send(groupid, photos, randlist)
	}
	log.Printf("%d/%d photos share to: %s\n", *shareN, num, *groupID)
}

func doShareToGroupByPhotoID() {
	var photoIDs []string
	for _, v := range strings.Split(*photoID, ",") {
		if len(v) > 0 {
			photoIDs = append(photoIDs, v)
		}
	}
	if len(photoIDs) == 0 {
		return
	}

	var groupIDs []string
	for _, v := range strings.Split(*groupID, ",") {
		if len(v) > 0 {
			groupIDs = append(groupIDs, v)
		}
	}
	if len(groupIDs) == 0 {
		return
	}

	wg.Add(len(photoIDs) * len(groupIDs))
	limit := make(chan struct{}, *concurrency)

	for _, groupid := range groupIDs {
		for _, photoid := range photoIDs {
			go addToPoolByPhotoID(groupid, photoid, limit)
		}
	}
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	f = flickr.NewFlickr(*apikey, *secret)
	f.AuthToken = os.Getenv("FLICKRUSERTOKEN")

	if *tags == "" {
		photos = fromSets()
	} else {
		photos = fromSearch()
	}

	doShareToGroup()
	doShareToGroupByPhotoID()
	wg.Wait()
}
