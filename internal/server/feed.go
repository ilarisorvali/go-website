package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/feeds"
)

const url = "https://sorvali.net/posts/"

func (app *application) generateAtomFeed() string {
	now := time.Now()
	feed := &feeds.Feed{
		Title:   "sorvali.net posts",
		Link:    &feeds.Link{Href: "http://sorvali.net/posts"},
		Author:  &feeds.Author{Name: "Ilari Sorvali"},
		Created: now,
	}

	fmt.Println("asdasdasdasd")

	for _, item := range app.contentCache.Blog {
		fmt.Println(item.Meta.Date)
		link := url + item.Meta.Slug

		entry := &feeds.Item{
			Title:       item.Meta.Title,
			Link:        &feeds.Link{Href: link},
			Description: item.Meta.Description,
			Created:     item.Meta.Date.Time,
		}

		feed.Add(entry)

	}

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(atom)
	return atom
}
