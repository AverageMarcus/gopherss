package server

import (
	"github.com/averagemarcus/gopherss/internal/feeds"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	FeedStore *feeds.FeedStore
}

func (a *API) GetFeeds(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetFeeds())
}

func (a *API) GetFeed(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetFeed(c.Params("id")))
}

func (a *API) DeleteFeed(c *fiber.Ctx) error {
	a.FeedStore.DeleteFeed(c.Params("id"))
	return nil
}

func (a *API) PostFeed(c *fiber.Ctx) error {
	url := ""
	if err := c.BodyParser(&url); err != nil {
		return err
	}
	return c.JSON(feeds.RefreshFeed(url))
}

func (a *API) GetItem(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetItem(c.Params("id")))
}

func (a *API) GetUnread(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetUnread())
}

func (a *API) GetSaved(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetSaved())
}

func (a *API) GetAll(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetAll())
}

func (a *API) PostRead(c *fiber.Ctx) error {
	a.FeedStore.MarkAsRead(c.Params("id"))
	return nil
}

func (a *API) PostReadAll(c *fiber.Ctx) error {
	ids := &[]string{}
	if err := c.BodyParser(ids); err != nil {
		return err
	}
	for _, id := range *ids {
		a.FeedStore.MarkAsRead(id)
	}
	return c.JSON(a.FeedStore.GetUnread())
}

func (a *API) RefreshAll(c *fiber.Ctx) error {
	for _, feed := range *a.FeedStore.GetFeeds() {
		feeds.RefreshFeed(feed.FeedURL)
	}

	return c.JSON(a.FeedStore.GetUnread())
}

func (a *API) SaveItem(c *fiber.Ctx) error {
	a.FeedStore.ToggleSaved(c.Params("id"))
	return nil
}
