package flickr

import (
	"encoding/json"
	"log"

	"github.com/toomore/lazyflickrgo/jsonstruct"
	"github.com/toomore/lazyflickrgo/utils"
)

// peopleFindBy username, email for find people
func (f Flickr) peopleFindBy(username, email string) jsonstruct.PeopleFindBy {
	data := make(map[string]string)

	if username != "" {
		data["method"] = "flickr.people.findByUsername"
		data["username"] = username
	} else if email != "" {
		data["method"] = "flickr.people.findByEmail"
		data["find_email"] = email
	} else {
		return jsonstruct.PeopleFindBy{}
	}

	jsonData := f.HTTPGet(utils.APIURL, data)

	var result jsonstruct.PeopleFindBy
	if err := json.Unmarshal(jsonData, &result); err != nil {
		log.Println(err)
	}
	return result
}

// PeopleFindByEmail for find user by email
func (f Flickr) PeopleFindByEmail(email string) jsonstruct.PeopleFindBy {
	return f.peopleFindBy("", email)
}

// PeopleFindByUsername for find user by username
func (f Flickr) PeopleFindByUsername(username string) jsonstruct.PeopleFindBy {
	return f.peopleFindBy(username, "")
}
