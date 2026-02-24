package tui

import (
	"fmt"

	"devbook_meeting2/code/pkg/devbook"
)

func printItems(store devbook.Store) {
	for _, item := range store.List() {
		fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)
	}
}

func addItemView(store devbook.Store) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	description, err := AskForString("Description")
	if err != nil {
		return err
	}

	url, err := AskForString("URL")
	if err != nil {
		return err
	}

	store.Add(devbook.Item{
		Name:        name,
		Description: description,
		URL:         url,
	})
	return nil
}

func deleteItemView(store devbook.Store) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	store.Delete(name)
	return nil
}

func MenuView(store devbook.Store) error {
	choices := map[string]string{
		"a": "Add",
		//"e":   "Edit",
		"d": "Delete",
		"l": "List",
		"q": "Exit",
	}

	for {
		choice, err := AskForMapChoice("Choose:", choices)
		if err != nil {
			return err
		}

		switch choice {
		case "a":
			if err = addItemView(store); err != nil {
				return err
			}
		case "e":
			fmt.Println("Not yet implemented")
		case "d":
			if err = deleteItemView(store); err != nil {
				return err
			}
		case "l":
			printItems(store)
		case "q":
			return nil
		}
		fmt.Println()
	}
}
