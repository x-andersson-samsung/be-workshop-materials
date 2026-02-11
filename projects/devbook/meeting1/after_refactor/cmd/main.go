package main

import (
	"devbook_meeting1/after_refactor/pkg/devbook"
	"devbook_meeting1/after_refactor/pkg/tui"
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
