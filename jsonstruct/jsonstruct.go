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
	ID         string `json:"id"`
	Author     string `json:"author"`
	Authorname string `json:"authorname"`
	Content    string `json:"_content"`
	//MachineTag bool   `json:"machine_tag"`
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

// Place struct
type Place struct {
	Content string `json:"_content"`
	PlaceID string `json:"place_id"`
	Woeid   string `json:"woeid"`
}

// Location struct
type Location struct {
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Accuracy      string `json:"accuracy"`
	Context       string `json:"context"`
	PlaceID       string `json:"place_id"`
	Woeid         string `json:"woeid"`
	Neighbourhood Place  `json:"neighbourhood"`
	Locality      Place  `json:"locality"`
	County        Place  `json:"county"`
	Region        Place  `json:"region"`
	Country       Place  `json:"country"`
}

// PhotosGetInfo in flickr.photos.getInfo
type PhotosGetInfo struct {
	Photo struct {
		ID           string   `json:"id"`
		Dateuploaded string   `json:"dateuploaded"`
		License      string   `json:"license"`
		Media        string   `json:"media"`
		Orgformat    string   `json:"originalformat"`
		Orgsecret    string   `json:"originalsecret"`
		Secret       string   `json:"secret"`
		Server       string   `json:"server"`
		Views        string   `json:"views"`
		Farm         int64    `json:"farm"`
		Rotation     int64    `json:"rotation"`
		Comments     Content  `json:"comments"`
		Description  Content  `json:"description"`
		Title        Content  `json:"title"`
		Tags         Tags     `json:"tags"`
		Urls         URL      `json:"urls"`
		Location     Location `json:"location"`
		Owner        struct {
			Iconfarm   int64  `json:"iconfarm"`
			Iconserver string `json:"iconserver"`
			Location   string `json:"location"`
			Nsid       string `json:"nsid"`
			PathAlias  string `json:"path_alias"`
			Realname   string `json:"realname"`
			Username   string `json:"username"`
		} `json:"owner"`
		Dates struct {
			Posted           string `json:"posted"`
			Taken            string `json:"taken"`
			Takengranularity int64  `json:"takengranularity"`
			Takenunknown     string `json:"takenunknown"`
			Lastupdate       string `json:"lastupdate"`
		} `json:"dates"`
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

// License struct
type License struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PhotosLicenses struct
type PhotosLicenses struct {
	Licenses struct {
		License []License `json:"license"`
	} `json:"licenses"`
	Common
}

// PhotosSearch in flickr.photos.search
type PhotosSearch struct {
	Photos struct {
		Page    int     `json:"page"`
		Pages   int     `json:"pages"`
		Perpage int     `json:"perpage"`
		Total   string  `json:"total"`
		Photo   []Photo `json:"photo"`
	} `json:"photos"`
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
	Lang            string  `json:"lang"`
	Ispoolmoderated int     `json:"ispoolmoderated"`
	IconFarm        int     `json:"iconfarm"`
	Name            Content `json:"name"`
	Description     Content `json:"description"`
	Members         Content `json:"members"`
	Poolcount       Content `json:"pool_count"`
	Topiccount      Content `json:"topic_count"`
	Privacy         Content `json:"privacy"`
	Blast           struct {
		Content        string `json:"_content"`
		DateBlastAdded string `json:"date_blast_added"`
		UserID         string `json:"user_id"`
	} `json:"blast"`
	Throttle struct {
		Count     int64  `json:"count,string"`
		Mode      string `json:"mode"`
		Remaining int64  `json:"remaining"`
	} `json:"throttle"`
	//Rules           Content `json:"rules"`
}

// GroupsGetInfo struct
type GroupsGetInfo struct {
	Group Group `json:"group"`
	Common
}

// PeopleFindBy struct
type PeopleFindBy struct {
	User struct {
		ID       string `json:"id"`
		Nsid     string `json:"nsid"`
		Username struct {
			Content string `json:"_content"`
		} `json:"username"`
	} `json:"user"`
	Stat    string `json:"stat"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// PeopleGetGroups struct
type PeopleGetGroups struct {
	Groups struct {
		Group []PeopleGroup `json:"group"`
	} `json:"groups"`
	Stat string `json:"stat"`
}

// PeopleGroup struct
type PeopleGroup struct {
	Nsid           string `json:"nsid"`
	Name           string `json:"name"`
	Iconfarm       int    `json:"iconfarm"`
	Iconserver     string `json:"iconserver"`
	Admin          int    `json:"admin"`
	Eighteenplus   int    `json:"eighteenplus"`
	InvitationOnly int    `json:"invitation_only"`
	Members        string `json:"members"`
	PoolCount      string `json:"pool_count"`
	IsMember       int    `json:"is_member"`
	IsModerator    int    `json:"is_moderator"`
	IsAdmin        int    `json:"is_admin"`
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
