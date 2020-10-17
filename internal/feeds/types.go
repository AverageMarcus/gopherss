package feeds

import (
	"encoding/xml"
	"time"

	"gorm.io/gorm"
)

type FeedStore struct {
	db *gorm.DB
}

type Feed struct {
	ID          string `gorm:"primaryKey"` // Base64 of FeedURL
	Title       string
	Description string
	HomepageURL string
	FeedURL     string
	ImageURL    string
	LastUpdated *time.Time
	Items       []Item `gorm:"foreignKey:FeedID"`
	UnreadCount int
}

func (feed *Feed) AfterFind(tx *gorm.DB) (err error) {
	feed.UnreadCount = 0
	for _, item := range feed.Items {
		if !item.Read {
			feed.UnreadCount++
		}
	}
	return
}

type Item struct {
	ID          string `gorm:"primaryKey"` // Base64 of GUID
	Title       string
	Description string
	Content     string
	URL         string
	ImageURL    string
	LastUpdated *time.Time
	Created     *time.Time
	GUID        string
	FeedID      string
	Read        bool
	Save        bool
}

type ItemWithFeed struct {
	Item

	FeedTitle       string
	FeedHomepageURL string
}

type Opml struct {
	XMLName  xml.Name  `xml:"opml"`
	Version  string    `xml:"version,attr"`
	Outlines []Outline `xml:"body>outline"`
}

type Outline struct {
	Title  string `xml:"title,attr"`
	XmlUrl string `xml:"xmlUrl,attr"`
}
