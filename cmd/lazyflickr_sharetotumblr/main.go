package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/toomore/lazyflickrgo/flickr"
	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazytumblr/tumblr"
)

// https://farm{farm-id}.staticflickr.com/{server-id}/{id}_{o-secret}_o.(jpg|gif|png)
const imageFormat = "https://farm%d.staticflickr.com/%s/%s_%s_o.%s"

var (
	apikey = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	info   = color.New(color.Bold, color.FgGreen).SprintfFunc()
	warn   = color.New(color.Bold, color.FgRed).SprintfFunc()
	wg     sync.WaitGroup
)

func main() {
	flag.Parse()
	f := flickr.NewFlickr(*apikey, *secret)
	t := tumblr.NewTumblr(
		os.Getenv("TUMBLRCONSUMERKEY"), os.Getenv("TUMBLRCONSUMERSECRET"))
	t.BaseHost = os.Getenv("TUMBLRUSERBASEHOST")
	t.Token = os.Getenv("TUMBLRUSERTOKEN")
	t.TokenSecret = os.Getenv("TUMBLRUSERSECRET")

	num := len(flag.Args())
	photolist := make([]jsonstruct.PhotosGetInfo, num)
	wg.Add(num)
	for i, photoID := range flag.Args() {
		go func(i int, photoID string) {
			runtime.Gosched()
			defer wg.Done()
			photolist[i] = f.PhotosGetInfo(photoID)
		}(i, photoID)
	}

	wg.Wait()

	for _, photo := range photolist {
		photoURL := fmt.Sprintf(imageFormat,
			photo.Photo.Farm,
			photo.Photo.Server,
			photo.Photo.ID,
			photo.Photo.Orgsecret,
			photo.Photo.Orgformat)

		tagslist := make([]string, len(photo.Photo.Tags.Tag))
		for i, val := range photo.Photo.Tags.Tag {
			tagslist[i] = val.Raw
		}
		//tags := strings.Join(tagslist, ",")

		//log.Printf("%+v\n", photo)
		//log.Println(photoURL)
		//log.Println(tags)
		//log.Println(strings.Replace(photo.Photo.Description.Content, "\n", "<br>", -1))
		//log.Println(photo.Photo.Urls.URL[0].Content)

		args := make(map[string]string)
		args["source"] = photoURL
		args["tags"] = strings.Join(tagslist, ",")
		args["caption"] = fmt.Sprintf("<h2>%s</h2><p>%s</p>",
			photo.Photo.Title.Content,
			strings.Replace(photo.Photo.Description.Content, "\n", "<br>", -1),
		)
		args["source_url"] = photo.Photo.Urls.URL[0].Content

		resp := t.PostPhoto(args, nil)
		if resp.StatusCode == 201 {
			log.Println(info("[%s] %s", photo.Photo.ID, photo.Photo.Title.Content))
		} else {
			log.Println(warn("[Error:%d] [%s] %s", resp.StatusCode, photo.Photo.ID, photo.Photo.Title.Content))
		}
	}
}
