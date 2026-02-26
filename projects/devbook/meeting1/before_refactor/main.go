package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const prompt = "> "

func readLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), scanner.Err()
}

func AskForString(question string) (string, error) {
	_, err := fmt.Printf("%s\n%s", question, prompt)
	if err != nil {
		return "", err
	}

	answer, err := readLine()
	return strings.TrimSpace(answer), err
}

func AskForMapChoice(question string, choices map[string]string) (string, error) {
	// Print question
	_, err := fmt.Println(question)
	if err != nil {
		return "", err
	}

	// Check for longest key for formatting
	maxLength := 0
	for key := range choices {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}

	// Sort by key
	type choiceStruct struct {
		key   string
		value string
	}
	sortedChoices := make([]choiceStruct, 0, len(choices))
	for k, v := range choices {
		sortedChoices = append(sortedChoices, choiceStruct{key: k, value: v})
	}
	sort.Slice(sortedChoices, func(i, j int) bool { return sortedChoices[i].key < sortedChoices[j].key })

	// Print choices
	for _, choice := range sortedChoices {
		if _, err = fmt.Printf("%*s: %s\n", maxLength, choice.key, choice.value); err != nil {
			return "", err
		}
	}

	// Print prompt
	if _, err = fmt.Print(prompt); err != nil {
		return "", err
	}

	// Read answer
	answer, err := readLine()
	if err != nil {
		return "", err
	}

	// Validate
	if _, ok := choices[answer]; !ok {
		fmt.Println("Invalid choice")
		return AskForMapChoice(question, choices)
	}

	return answer, nil
}

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

func printItems() {
	for _, item := range store {
		fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)
	}
}

func addItemView() error {
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

	store[name] = Item{
		Name:        name,
		Description: description,
		URL:         url,
	}
	return nil
}

func editItemView() error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	if _, ok := store[name]; !ok {
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
		delete(store, name)
	}

	store[name] = Item{
		Name:        newName,
		Description: description,
		URL:         url,
	}
	return nil
}

func deleteItemView() error {
	name, err := AskForString("Name")
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

func menuView() error {
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
			if err = addItemView(); err != nil {
				return err
			}
		case "e":
			if err = editItemView(); err != nil {
				return err
			}
		case "d":
			if err = deleteItemView(); err != nil {
				return err
			}
		case "l":
			printItems()
		case "q":
			return nil
		}
		fmt.Println()
	}
}

func main() {
	if err := menuView(); err != nil {
		panic(err)
	}
}
