package mApp

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

func (ma *MApp) generateRSS() string {
	layout := "2006-01-02 15:04:05"
	now := time.Now()

	feed := &feeds.Feed{
		Created:     now,
		Title:       ma.Config.MSite.Info.Title,
		Link:        &feeds.Link{Href: ma.Config.MSite.Info.Link},
		Author:      &feeds.Author{Name: ma.Config.MSite.Info.Author, Email: ma.Config.MSite.Info.Email},
		Description: ma.Config.MSite.Info.Description,
	}

	count := 0
	var items []*feeds.Item
	for _, post := range ma.Posts {
		if count >= ma.Config.MSite.Post.RecentPost.Number {
			break
		}
		
		updated, err := time.Parse(layout, post.Date)
		if err != nil {
			log.Println("Error parsing updated time ", post.Date)
			continue
		}

		item := &feeds.Item{
			Title:   post.Title,
			Id:      strconv.FormatUint(post.Index, 10),
			Link:    &feeds.Link{Href: fmt.Sprintf("%s/post/%s", ma.Config.MSite.Info.Link, post.HtmlHash)},
			Author:  &feeds.Author{Name: ma.Config.MSite.Info.Author, Email: ma.Config.MSite.Info.Email},
			Updated: updated,
		}

		items = append(items, item)
		count++
	}

	feed.Items = items
	rss, err := feed.ToRss()
	if err != nil {
		log.Println("Error generating rss ", err)
	}

	return rss
}
