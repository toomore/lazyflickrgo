LazyFlickrGo
============

[![GoDoc](https://godoc.org/github.com/toomore/lazyflickrgo?status.svg)](https://godoc.org/github.com/toomore/lazyflickrgo)

Because of more photos need to share, but no time to do that. OK, write some code for sharing photos.

Completed:
- AuthGetFrob (`flickr.auth.getFrob`)
- AuthGetToken (`flickr.auth.getToken`)
- GroupsGetInfo (`flickr.groups.getInfo`)
- GroupsPoolsAdd (`flickr.groups.pools.add`)
- PeopleFindByEmail (`flickr.people.findByEmail`)
- PeopleFindByUsername (`flickr.people.findByUsername`)
- PeopleGetGroups(`flickr.people.getGroups`)
- PhotosGetInfo (`flickr.photos.getInfo`)
- PhotosSearch (`flickr.photos.search`)
- PhotosetsGetInfo (`flickr.photosets.getInfo`)
- PhotosetsGetPhotos (`flickr.photosets.getPhotos`)
  - PhotosetsGetPhotosAll *(Get all pages data)*

Environment vars *(some vars just for testing)*:
- `FLICKRAPIKEY` *(Get from [Apps.Create](https://www.flickr.com/services/apps/create/))*
- `FLICKRSECRET` *(Get from [Apps.Create](https://www.flickr.com/services/apps/create/))*
- `FLICKRUSERTOKEN` *(Get from user auth `AuthGetFrob`, `AuthGetToken`)*
- `FLICKRUSER` *(Only for testing)*
- `FLICKRGROUPID` *(Only for testing)*
- `FLICKRPHOTOID` *(Only for testing)*

... and still in development.
