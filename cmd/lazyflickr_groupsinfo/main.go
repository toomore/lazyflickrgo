// cmd/lazyflickr_groupsinfo search group id by name.
/*
Install:

	go install github.com/toomore/lazyflickrgo/cmd/lazyflickr_groupsinfo

Usage:

	lazyflickr_groupsinfo [flags] <group name>[ <group name> ...]

The flags are:

	-apikey
		Flickr API key, default get from env `FLICKRAPIKEY`

	-secret
		Flickr secret, default get from env `FLICKRSECRET`


Example:
	lazyflickr_groupsinfo lomo lomo.tw

Result:
	2016/07/24 20:08:36 [lomotw] LOMO.tw => 72262428@N00
	2016/07/24 20:08:36 [lomo] LOMO => 40732537997@N01


*/
package main

import (
	"flag"
	"fmt"
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
	user    = flag.String("u", "", "Get user's all groups info")
	info    = color.New(color.Bold, color.FgGreen).SprintfFunc()
	wg      sync.WaitGroup
	numChan = runtime.NumCPU() * 4
	f       *flickr.Flickr
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
	result := make(chan jsonstruct.GroupsGetInfo, numChan)

	go func() {
		for val := range name {
			go func(val string) {
				defer wg.Done()
				result <- f.GroupsGetInfo("", val, false)
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
	f = flickr.NewFlickr(*apikey, *secret)

	if *user == "" {
		inData := iniPut(flag.Args())
		result := dogetInfo(inData)
		<-outPut(result)
	} else {
		f.AuthToken = os.Getenv("FLICKRUSERTOKEN")
		userGroups := f.PeopleGetGroups(*user, "")
		for i, v := range userGroups.Groups.Group {
			fmt.Printf("\"%d\",\"%s\",\"%s\",\"%s\",\"%s\"\n", i, v.Name, v.Nsid, v.Members, v.PoolCount)
		}
	}

}
