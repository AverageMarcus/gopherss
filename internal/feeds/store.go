package feeds

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func (fs *FeedStore) getDB() *gorm.DB {
	if fs.db == nil {
		db, err := gorm.Open(sqlite.Open(viper.GetString("DB_PATH")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&Feed{})
		db.AutoMigrate(&Item{})

		fs.db = db
	}

	return fs.db
}

func (fs *FeedStore) GetFeed(id string) *Feed {
	feed := &Feed{}
	fs.getDB().Preload("Items").Where("id = ?", id).First(feed)
	return feed
}

func (fs *FeedStore) GetItem(id string) *Item {
	item := &Item{}
	fs.getDB().Where("id = ?", id).First(item)
	return item
}

func (fs *FeedStore) GetFeeds() *[]Feed {
	feeds := &[]Feed{}
	fs.getDB().Order("title asc").Find(feeds)
	return feeds
}

func (fs *FeedStore) GetUnread() *[]ItemWithFeed {
	items := &[]ItemWithFeed{}
	fs.getDB().Table("items").
		Where("read = ?", false).
		Select("items.*, feeds.title as feed_title, feeds.homepage_url as feed_homepage_url").
		Order("items.created desc, items.title").
		Joins("left join feeds on feeds.id = items.feed_id").
		Find(items)

	return items
}

func (fs *FeedStore) SaveFeed(feed Feed) {
	fs.getDB().Omit("Items").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "description", "homepage_url", "image_url", "last_updated"}),
	}).Create(feed)

	for _, item := range feed.Items {
		fs.getDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"title", "description", "content", "url", "image_url", "last_updated"}),
		}).Create(item)
	}
}

func (fs *FeedStore) MarkAsRead(itemID string) {
	item := &Item{}
	fs.getDB().Where("id = ?", itemID).First(item)

	item.Read = true

	fs.getDB().Save(*item)
}
