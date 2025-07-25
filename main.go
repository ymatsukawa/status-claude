package main

import (
	"fmt"

	c "github.com/ymatsukawa/sclaude/core"
	d "github.com/ymatsukawa/sclaude/decorator"
)

func main() {
	rss := c.NewRss()

	if err := rss.Parse(); err != nil {
		panic(err)
	}

	feed := rss.GetFeed()
	if feed == nil {
		panic("rss feed(s) not found")
	}
	status := c.NewAnthropicStatus(feed.Items)

	if !status.IsElevatedErrors() {
		fmt.Printf("\n%s\n\n", d.Colorize("on going.", d.YELLOW))
		return
	}

	fmt.Println(d.Colorize(status.GetErrorMessage(), d.RED))
}
