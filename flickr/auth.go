package flickr

import (
	"encoding/json"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// AuthGetFrob to get Frob link.
func (f Flickr) AuthGetFrob() jsonstruct.AuthGetFrob {
	Args := map[string]string{"method": "flickr.auth.getFrob"}
	jsonData := f.HTTPGet(utils.APIURL, Args)

	var data jsonstruct.AuthGetFrob
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}
	return data
}

// AuthGetToken to get user auth token.
func (f Flickr) AuthGetToken(frob string) jsonstruct.AuthGetToken {
	args := make(map[string]string)
	args["method"] = "flickr.auth.getToken"
	args["frob"] = frob

	jsonData := f.HTTPGet(utils.APIURL, args)
	log.Printf("%s\n", jsonData)

	var data jsonstruct.AuthGetToken
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}

	return data
}
