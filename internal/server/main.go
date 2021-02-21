package server

import (
	"fmt"
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/averagemarcus/gopherss/internal/feeds"
)

var api API

func init() {
	api = API{
		FeedStore: &feeds.FeedStore{},
	}
}

func Start(port string) error {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	engine.AddFunc("htmlSafe", func(html string) template.HTML {
		return template.HTML(html)
	})
	engine.AddFunc("humanDate", func(date time.Time) template.HTML {
		return template.HTML(humanize.Time(date))
	})
	engine.AddFunc("coalesce", func(args ...*string) string {
		for _, s := range args {
			if s != nil && *s != "" {
				return *s
			}
		}
		return ""
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views")
	app.Post("/opml", PostOPML)

	// API
	app.Get("/api/feeds", api.GetFeeds)
	app.Post("/api/feeds", api.PostFeed)
	app.Get("/api/feed/:id", api.GetFeed)
	app.Delete("/api/feed/:id", api.DeleteFeed)
	app.Get("/api/item/:id", api.GetItem)
	app.Post("/api/item/:id/save", api.SaveItem)
	app.Get("/api/unread", api.GetUnread)
	app.Get("/api/saved", api.GetSaved)
	app.Get("/api/all", api.GetAll)
	app.Post("/api/read/:id", api.PostRead)
	app.Post("/api/read", api.PostReadAll)
	app.Get("/api/refresh", api.RefreshAll)

	return app.Listen(fmt.Sprintf(":%s", port))
}

func PostOPML(c *fiber.Ctx) error {
	opml := &feeds.Opml{}
	if err := c.BodyParser(opml); err != nil {
		return err
	}

	f := []feeds.Feed{}
	for _, outline := range opml.Outlines {
		f = append(f, feeds.RefreshFeed(outline.XmlUrl))
	}

	return c.JSON(f)
}
