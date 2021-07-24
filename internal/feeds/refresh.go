package feeds

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

func RefreshFeed(feedUrl string) Feed {
	fmt.Printf("Refreshing %s\n", feedUrl)
	var feed Feed
	f, err := fp.ParseURL(feedUrl)
	if err != nil && err == gofeed.ErrFeedTypeNotDetected {
		foundFeed := loadFeedFromWebpage(feedUrl)
		if foundFeed != nil {
			feed = *foundFeed
		}
	} else if err != nil {
		fmt.Printf("Failed to refresh %s - %v\n", feedUrl, err)
	} else {
		imageURL := ""
		if f.Image != nil {
			imageURL = f.Image.URL
		}

		feed = Feed{
			ID:          strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(feedUrl)), "/", ""),
			Title:       f.Title,
			Description: f.Description,
			HomepageURL: f.Link,
			FeedURL:     feedUrl,
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

func loadFeedFromWebpage(webpageUrl string) *Feed {
	res, err := http.Get(webpageUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	feedUrl, ok := doc.Find(`[rel="alternate"][type="application/rss+xml"]`).First().Attr("href")
	if ok {
		if !strings.HasPrefix(feedUrl, "http") {
			parsedUrl, _ := url.Parse(webpageUrl)
			feedUrl = fmt.Sprintf("%s://%s%s", parsedUrl.Scheme, parsedUrl.Host, feedUrl)
		}

		feed := RefreshFeed(feedUrl)
		return &feed
	}

	return nil
}
