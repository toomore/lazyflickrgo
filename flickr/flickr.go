// Package flickr for api.
package flickr

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/toomore/lazyflickrgo/simplecache"
	"github.com/toomore/lazyflickrgo/utils"
)

// Flickr struct
type Flickr struct {
	args      map[string]string
	secretKey string
	AuthToken string
}

var tempDir string
var cache *simplecache.SimlpleCache

func init() {
	cache = simplecache.NewSimpleCache("", "lzf")
	log.Println("Temp Dir: ", cache.Dir, cache.Folder)
}

// NewFlickr is to new a request.
func NewFlickr(APIKey string, SecretKey string) *Flickr {
	args := make(map[string]string)

	// Default args.
	args["format"] = "json"
	args["nojsoncallback"] = "1"
	args["api_key"] = APIKey

	return &Flickr{
		args:      args,
		secretKey: SecretKey,
	}
}

// HTTPGet method request.
func (f Flickr) HTTPGet(URL string, Args map[string]string) []byte {
	for key, val := range f.args {
		Args[key] = val
	}

	if _, ok := Args["api_sig"]; ok {
		delete(Args, "api_sig")
	}
	Args["api_sig"] = utils.Sign(Args, f.secretKey)

	query := url.Values{}
	for key, val := range Args {
		query.Set(key, val)
	}

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatalln(err)
	}
	url.RawQuery = query.Encode()
	var data []byte
	if data, _ = cache.Get(Args["api_sig"]); data == nil {
		log.Println("Get: ", url.String())
		resp, err := http.Get(url.String())
		if err != nil {
			log.Fatalln(err)
		}

		data, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		cache.Set(Args["api_sig"], data)
	}

	return data
}

// HTTPPost method request.
func (f Flickr) HTTPPost(urlpath string, Data map[string]string) []byte {
	for key, val := range f.args {
		Data[key] = val
	}

	Data["api_sig"] = utils.Sign(Data, f.secretKey)

	log.Printf("Post: %+v %s", Data, urlpath)

	query := url.Values{}
	for key, val := range Data {
		query.Set(key, val)
	}

	resp, err := http.PostForm(urlpath, query)
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	return data
}
