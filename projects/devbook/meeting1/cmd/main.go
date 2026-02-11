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

func addItemView(cli *input.CLI) error {
	name, err := cli.AskForString("Name")
	if err != nil {
		return err
	}

	description, err := cli.AskForString("Description")
	if err != nil {
		return err
	}

	url, err := cli.AskForString("URL")
	if err != nil {
		return err
	}

	AddItem(name, description, url)
	return nil
}

func deleteItemView(cli *input.CLI) error {
	name, err := cli.AskForString("Name")
	if err != nil {
		return err
	}

	if _, ok := store[name]; !ok {
		fmt.Println("Item does not exist")
		return nil
	}

	delete(store, name)
	return nil
}

func optionView(cli *input.CLI) error {
	options := map[string]string{
		"a": "Add",
		//"e":   "Edit",
		"d": "Delete",
		"l": "List",
		"q": "Exit",
	}

	for {
		option, err := cli.AskForMapOption("Choose option", options)
		if err != nil {
			return err
		}

		switch option {
		case "a":
			if err = addItemView(cli); err != nil {
				return err
			}
		case "e":
			fmt.Println("Not yet implemented")
		case "d":
			if err = deleteItemView(cli); err != nil {
				return err
			}
		case "l":
			PrintItems(ListItems())
		case "q":
			return nil
		}
		fmt.Println()
	}
}

func main() {
	cli := input.NewCLI(os.Stdin, os.Stdout)
	if err := optionView(cli); err != nil {
		panic(err)
	}
}
