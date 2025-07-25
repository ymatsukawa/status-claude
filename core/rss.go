package core

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

const (
	STATUS_PAGE_RSS_URL = "https://status.anthropic.com/history.rss"
)

type IRss interface {
	Parse() error
	GetFeed() *gofeed.Feed
}

type Rss struct {
	Parser *gofeed.Parser
	Feed   *gofeed.Feed
}

func NewRss() IRss {
	return &Rss{
		Parser: gofeed.NewParser(),
		Feed:   nil,
	}
}

func (r *Rss) Parse() error {
	var err error
	r.Feed, err = r.Parser.ParseURL(STATUS_PAGE_RSS_URL)
	if err != nil {
		return fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return nil
}

func (r *Rss) GetFeed() *gofeed.Feed {
	if r.Feed == nil {
		return nil
	}
	return r.Feed
}
