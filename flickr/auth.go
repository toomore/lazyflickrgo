package flickr

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// AuthGetFrob to get Frob link.
func (f Flickr) AuthGetFrob() jsonstruct.AuthGetFrob {
	Args := map[string]string{"method": "flickr.auth.getFrob"}
	resp := f.HTTPGet(utils.APIURL, Args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

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

	resp := f.HTTPGet(utils.APIURL, args)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", jsonData)
	defer resp.Body.Close()

	var data jsonstruct.AuthGetToken
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println(err)
	}

	return data
}
