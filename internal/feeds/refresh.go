package feeds

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
)

var fp = gofeed.NewParser()
var feedStore = &FeedStore{}

func Refresh() error {
	interval := viper.GetInt("REFRESH_TIMEOUT")

	for {
		fmt.Println("Refreshing feeds...")

		for _, feed := range *feedStore.GetFeeds() {
			go RefreshFeed(feed.FeedURL)
		}

		fmt.Println("Reaping old items...")
		feedStore.DeleteOldReadItems()

		fmt.Printf("Going to sleep for %d minutes\n", interval)
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}

func RefreshFeed(url string) Feed {
	fmt.Printf("Refreshing %s\n", url)
	var feed Feed
	f, err := fp.ParseURL(url)
	if err != nil {
		fmt.Printf("Failed to refresh %s\n", url)
	} else {
		imageURL := ""
		if f.Image != nil {
			imageURL = f.Image.URL
		}

		feed = Feed{
			ID:          strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(url)), "/", ""),
			Title:       f.Title,
			Description: f.Description,
			HomepageURL: f.Link,
			FeedURL:     url,
			ImageURL:    imageURL,
			LastUpdated: f.UpdatedParsed,
			Items:       []Item{},
		}
		for _, item := range f.Items {
			imageURL := ""
			if f.Image != nil {
				imageURL = f.Image.URL
			}

			createdTime := item.PublishedParsed
			if createdTime == nil {
				createdTime = item.UpdatedParsed
			}

			feed.Items = append(feed.Items, Item{
				ID:          strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(item.GUID)), "/", ""),
				Title:       item.Title,
				Description: item.Description,
				Content:     item.Content,
				URL:         item.Link,
				ImageURL:    imageURL,
				LastUpdated: item.UpdatedParsed,
				Created:     createdTime,
				GUID:        item.GUID,
				FeedID:      feed.ID,
			})
		}
		feedStore.SaveFeed(feed)

		fmt.Printf("Finished refreshing '%s'\n", feed.Title)
	}

	return feed
}
