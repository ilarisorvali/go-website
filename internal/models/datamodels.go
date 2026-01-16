package models

import (
	"html/template"
	"time"
)

type ContentItem struct {
	Title       string    //meta: title of post
	Description string    //meta: description of post
	Slug        string    //meta: url slug of post
	Draft       bool      //is the ContentItem a draft, ie. is it visible
	Date        time.Time //timestamp of post, not a real 'hard' timestamp

	Content template.HTML //parsed main markdown content

}

type Pages struct {
	Posts []ContentItem
}
