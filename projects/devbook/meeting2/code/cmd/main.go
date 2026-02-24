package main

import (
	"devbook_meeting2/code/pkg/devbook"
	"devbook_meeting2/code/pkg/tui"
)

func main() {
	store := make(devbook.Store)
	store.Add(devbook.Item{
		Name:        "item1",
		Description: "search engine",
		URL:         "https://google.pl",
	})
	store.Add(devbook.Item{
		Name:        "item2",
		Description: "other search engine",
		URL:         "https://bing.com",
	})

	if err := tui.MenuView(store); err != nil {
		panic(err)
	}
}
