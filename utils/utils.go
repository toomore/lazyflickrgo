// Package utils more tools.
package utils

import (
	"crypto/md5"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	// APIURL Flickr API
	APIURL = "https://api.flickr.com/services/rest/"
	// AUTHURL Flickr Auth URL
	AUTHURL = "http://flickr.com/services/auth/"
)

// Sign for API `api_sig`
func Sign(args map[string]string) string {
	keySortedList := make([]string, len(args))
	var loop int64
	for key := range args {
		keySortedList[loop] = key
		loop++
	}
	sort.Strings(keySortedList)
	hashList := make([]string, len(args)*2)
	for i, val := range keySortedList {
		hashList[2*i] = val
		hashList[2*i+1] = args[val]
	}
	hashstring := fmt.Sprintf("%s%s", os.Getenv("FLICKRSECRET"), strings.Join(hashList, ""))
	return fmt.Sprintf("%x", md5.Sum([]byte(hashstring)))
}
