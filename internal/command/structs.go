package command

import (
	"errors"
	"html"

	"github.com/EndocrineMSc/gator/internal/config"
	"github.com/EndocrineMSc/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmap map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if _, exists := c.Cmap[cmd.Name]; exists {
		return c.Cmap[cmd.Name](s, cmd)
	}
	return errors.New("command not found")
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmap[name] = f
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (r *RSSFeed) Unescape() {
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	r.Channel.Title = html.UnescapeString(r.Channel.Title)

	for _, i := range r.Channel.Item {
		i.Description = html.UnescapeString(i.Description)
		i.Title = html.UnescapeString(i.Title)
	}
}
