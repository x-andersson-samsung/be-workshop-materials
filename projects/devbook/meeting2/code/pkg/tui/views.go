package tui

import (
	"fmt"

	"devbook_meeting2/code/pkg/devbook"
)

type ItemStore interface {
	List() []devbook.Item
	Add(item devbook.Item)
	GetByName(name string) (devbook.Item, bool)
	DeleteByName(name string)
}

func printItems(store ItemStore) {
	for _, item := range store.List() {
		fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)
	}
}

func addItemView(store ItemStore) error {
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

func editItemView(store ItemStore) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	if _, ok := store.GetByName(name); !ok {
		fmt.Println("Item does not exist")
		return nil
	}

	newName, err := AskForString("New Name")
	if err != nil {
		return err
	}

	description, err := AskForString("New Description")
	if err != nil {
		return err
	}

	url, err := AskForString("New URL")
	if err != nil {
		return err
	}

	if newName != name {
		store.DeleteByName(name)
	}

	store.Add(devbook.Item{
		Name:        newName,
		Description: description,
		URL:         url,
	})
	return nil
}

func deleteItemView(store ItemStore) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	store.DeleteByName(name)
	return nil
}

func MenuView(store ItemStore) error {
	choices := map[string]string{
		"a": "Add",
		"e": "Edit",
		"d": "DeleteByID",
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
			if err = editItemView(store); err != nil {
				return err
			}
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
