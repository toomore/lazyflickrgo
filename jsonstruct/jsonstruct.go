// Package jsonstruct struct all flickr api json.
package jsonstruct

import (
	"net/url"
	"os"

	"github.com/toomore/lazyflickrgo/utils"
)

// Common return format.
type Common struct {
	Stat    string `json:"stat"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Photo in flickr.photos.search
type Photo struct {
	ID       string `json:"id"`
	Owner    string `json:"owner"`
	Title    string `json:"title"`
	Secret   string `json:"secret"`
	Server   string `json:"server"`
	Farm     int64  `json:"farm"`
	Ispublic int64  `json:"ispublic"`
	Isfriend int64  `json:"isfriend"`
	Isfamily int64  `json:"isfamily"`
}

// Photos in flickr.photos.search
type Photos struct {
	Page    int64   `json:"page"`
	Pages   int64   `json:"pages"`
	Perpage int64   `json:"perpage"`
	Total   string  `json:"total"`
	Photo   []Photo `json:"photo"`
}

// PhotosSearch in flickr.photos.search
type PhotosSearch struct {
	Photos Photos `json:"photos"`
	//Stat   string `json:"stat"`
	Common
}

// AuthGetFrob in flickr.auth.getfrob
type AuthGetFrob struct {
	Frob Content `json:"frob"`
	Common
}

// Auth struct
type Auth struct {
	Token Content `json:"token"`
	Perms Content `json:"perms"`
	User  User    `json:"user"`
}

// AuthGetToken struct
type AuthGetToken struct {
	Auth Auth `json:"auth"`
	Common
}

// User struct
type User struct {
	Nsid     string `json:"nsid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

// Content common content
type Content struct {
	Content string `json:"_content"`
}

// GetTokenURL to output link.
func (auth AuthGetFrob) GetTokenURL() string {
	args := make(map[string]string)
	args["api_key"] = os.Getenv("FLICKRAPIKEY")
	args["perms"] = "write"
	args["frob"] = auth.Frob.Content

	Args := url.Values{}
	for key, val := range args {
		Args.Set(key, val)
	}
	Args.Set("api_sig", utils.Sign(args))

	url, _ := url.Parse(utils.AUTHURL)
	url.RawQuery = Args.Encode()

	return url.String()
}
