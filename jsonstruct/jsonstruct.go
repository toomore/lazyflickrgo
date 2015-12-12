// Package jsonstruct struct all flickr api json.
package jsonstruct

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
	Stat   string `json:"stat"`
}
