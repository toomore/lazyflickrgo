package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/fatih/color"
	"github.com/toomore/lazyflickrgo/flickr"
)

var (
	apikey = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	info   = color.New(color.Bold, color.FgGreen).SprintfFunc()
	wg     sync.WaitGroup
)

func main() {
	flag.Parse()
	f := flickr.NewFlickr(*apikey, *secret)
	wg.Add(len(flag.Args()))
	for _, val := range flag.Args() {
		go func(val string) {
			runtime.Gosched()
			groupInfo := f.GroupsGetInfo("", val)
			log.Printf("%s => %s\n", val, info(groupInfo.Group.ID))
			wg.Done()
		}(val)
	}
	wg.Wait()
}
