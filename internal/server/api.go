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

func (a *API) GetItem(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetItem(c.Params("id")))
}

func (a *API) GetUnread(c *fiber.Ctx) error {
	return c.JSON(a.FeedStore.GetUnread())
}

func (a *API) PostRead(c *fiber.Ctx) error {
	a.FeedStore.MarkAsRead(c.Params("id"))
	return nil
}
