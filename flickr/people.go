package flickr

import (
	"encoding/json"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// PeopleFindByEmail for find people
func (f Flickr) PeopleFindByEmail(email string) jsonstruct.PeopleFindByEmail {
	data := make(map[string]string)
	data["method"] = "flickr.people.findByEmail"
	data["find_email"] = email

	jsonData := f.HTTPGet(utils.APIURL, data)

	var result jsonstruct.PeopleFindByEmail
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}
