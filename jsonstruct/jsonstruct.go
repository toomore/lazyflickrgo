// Package jsonstruct struct all flickr api json.
package jsonstruct

import (
	"net/url"

	"github.com/toomore/lazyflickrgo/utils"
)

// Common return format.
type Common struct {
	Stat    string `json:"stat"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Tags struct
type Tags struct {
	Tag []tag `json:"tag"`
}

type tag struct {
	Raw string `json:"raw"`
}

// URL struct
type URL struct {
	URL []urlstr `json:"url"`
}

type urlstr struct {
	Type    string `json:"type"`
	Content string `json:"_content"`
}

// PhotosGetInfo in flickr.photos.getInfo
type PhotosGetInfo struct {
	Photo struct {
		ID          string  `json:"id"`
		Secret      string  `json:"secret"`
		Orgsecret   string  `json:"originalsecret"`
		Orgformat   string  `json:"originalformat"`
		Server      string  `json:"server"`
		Farm        int64   `json:"farm"`
		Title       Content `json:"title"`
		Description Content `json:"description"`
		Tags        Tags    `json:"tags"`
		Urls        URL     `json:"urls"`
	} `json:"photo"`
	Common
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
	Page    string  `json:"page"`
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

// Photoset struct
type Photoset struct {
	ID        string `json:"id"`
	Primary   string `json:"primary"`
	Owner     string `json:"owner"`
	Ownername string `json:"ownername"`
	Title     string `json:"title"`
	Photos
}

// PhotosetInfo struct
type PhotosetInfo struct {
	ID               string  `json:"id"`
	Primary          string  `json:"primary"`
	Owner            string  `json:"owner"`
	Username         string  `json:"username"`
	Title            Content `json:"title"`
	Description      Content `json:"description"`
	Secret           string  `json:"secret"`
	Server           string  `json:"server"`
	Farm             int     `json:"farm"`
	Photos           int     `json:"photos"`
	CountViews       string  `json:"count_views"`
	CountComment     string  `json:"count_comments"`
	CountPhotos      string  `json:"count_photos"`
	CountVideos      int     `json:"count_videos"`
	CanComment       int     `json:"can_comment"`
	DateCreate       string  `json:"date_create"`
	DateUpdate       string  `json:"date_update"`
	CoverPhotoServer string  `json:"coverphoto_server"`
	CoverPhotoFarm   int     `json:"coverphoto_farm"`
}

// PhotosetsGetPhotos struct
type PhotosetsGetPhotos struct {
	Photoset Photoset `json:"photoset"`
	Common
}

// PhotosetsGetInfo struct
type PhotosetsGetInfo struct {
	Photoset PhotosetInfo `json:"photoset"`
	//Common
}

// Group struct
type Group struct {
	ID              string  `json:"id"`
	PathAlias       string  `json:"path_alias"`
	IconServer      string  `json:"iconserver"`
	IconFarm        int     `json:"iconfarm"`
	Name            Content `json:"name"`
	Description     Content `json:"description"`
	Rules           Content `json:"rules"`
	Members         Content `json:"members"`
	Poolcount       Content `json:"pool_count"`
	Topiccount      Content `json:"topic_count"`
	Privacy         Content `json:"privacy"`
	Lang            string  `json:"lang"`
	Ispoolmoderated int     `json:"ispoolmoderated"`
}

// GroupsGetInfo struct
type GroupsGetInfo struct {
	Group Group `json:"group"`
	Common
}

// GetTokenURL to output link.
func (auth AuthGetFrob) GetTokenURL(APIKey string, secretKey string) string {
	args := make(map[string]string)
	args["api_key"] = APIKey
	args["perms"] = "write"
	args["frob"] = auth.Frob.Content

	Args := url.Values{}
	for key, val := range args {
		Args.Set(key, val)
	}
	Args.Set("api_sig", utils.Sign(args, secretKey))

	url, _ := url.Parse(utils.AUTHURL)
	url.RawQuery = Args.Encode()

	return url.String()
}
