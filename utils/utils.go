package utils

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
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
	hashstring := strings.Join(hashList, "")
	return fmt.Sprintf("%x", md5.Sum([]byte(hashstring)))
}
