package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/fatih/color"
	"github.com/toomore/lazyflickrgo/flickr"
	"github.com/toomore/lazyflickrgo/jsonstruct"
)

var (
	apikey  = flag.String("apikey", os.Getenv("FLICKRAPIKEY"), "Flickr API Key")
	secret  = flag.String("secret", os.Getenv("FLICKRSECRET"), "Flickr secret")
	info    = color.New(color.Bold, color.FgGreen).SprintfFunc()
	wg      sync.WaitGroup
	numChan = runtime.NumCPU() * 4
)

func iniPut(args []string) <-chan string {
	result := make(chan string, numChan)
	wg.Add(len(args))

	go func() {
		defer close(result)
		for _, val := range args {
			result <- val
		}
	}()
	return result
}

func dogetInfo(name <-chan string) <-chan jsonstruct.GroupsGetInfo {
	f := flickr.NewFlickr(*apikey, *secret)
	result := make(chan jsonstruct.GroupsGetInfo, numChan)

	go func() {
		for val := range name {
			go func(val string) {
				defer wg.Done()
				result <- f.GroupsGetInfo("", val)
			}(val)
		}
	}()
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func outPut(result <-chan jsonstruct.GroupsGetInfo) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for data := range result {
			log.Printf("[%s] %s => %s\n",
				data.Group.PathAlias, data.Group.Name.Content, info(data.Group.ID))
		}
	}()
	return done
}

func main() {
	flag.Parse()

	inData := iniPut(flag.Args())
	result := dogetInfo(inData)
	<-outPut(result)
}
