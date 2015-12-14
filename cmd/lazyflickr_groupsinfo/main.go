package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/toomore/lazyflickrgo/flickr"
)

var (
	apikey = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
)

func main() {
	flag.Parse()
	f := flickr.NewFlickr(*apikey, *secret)
	var wg sync.WaitGroup
	wg.Add(len(flag.Args()))
	for _, val := range flag.Args() {
		go func(val string) {
			groupInfo := f.GroupsGetInfo("", val)
			log.Printf("%s: %s", val, groupInfo.Group.ID)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
