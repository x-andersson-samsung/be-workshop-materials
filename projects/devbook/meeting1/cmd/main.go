package main

import (
	"fmt"
	"os"

	"devbook_meeting1/pkg/input"
)

type Item struct {
	Name        string
	Description string

	URL string
}

var store = map[string]Item{
	"item1": {
		Name:        "item1",
		Description: "search engine",
		URL:         "https://google.pl",
	},
	"item2": {
		Name:        "item2",
		Description: "other search engine",
		URL:         "https://bing.com",
	},
}

func AddItem(name string, description string, url string) {
	store[name] = Item{
		Name:        name,
		Description: description,
		URL:         url,
	}
}

func ListItems() []Item {
	arr := make([]Item, 0, len(store))
	for _, item := range store {
		arr = append(arr, item)
	}
	return arr
}

func PrintItems(items []Item) {
	for _, item := range items {
		fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)
	}
}

func main() {
	cli := input.NewCLI(os.Stdin, os.Stdout)
	opt, err := cli.AskForOption("Which to remove?", []string{"item1", "item2"})
	if err != nil {
		fmt.Println("Invalid option")
	}
	fmt.Println(opt)
}
