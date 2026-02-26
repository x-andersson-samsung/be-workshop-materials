package devbook

import (
	"slices"
)

type Store struct {
	nextID int
	data   map[int]Item
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		data:   make(map[int]Item),
	}
}

func (store *Store) GetByID(id int) (Item, bool) {
	item, ok := store.data[id]
	return item, ok
}

func (store *Store) GetByName(name string) (Item, bool) {
	for _, item := range store.data {
		if item.Name == name {
			return item, true
		}
	}

	return Item{}, false
}

func (store *Store) Add(item Item) {
	item.ID = store.nextID
	store.nextID++

	store.data[item.ID] = item
}

func (store *Store) DeleteByID(id int) {
	delete(store.data, id)
}

func (store *Store) DeleteByName(name string) {
	for id, item := range store.data {
		if item.Name == name {
			delete(store.data, id)
			return
		}
	}
}

func (store *Store) List() []Item {
	// Expressive way
	arr := make([]Item, 0, len(store.data))
	for _, item := range store.data {
		arr = append(arr, item)
	}

	slices.SortFunc(arr, func(i, j Item) int {
		if i.ID == j.ID {
			return 0
		}
		if i.ID > j.ID {
			return 1
		}
		return -1
	})

	return arr

	// Compact way
	//return slices.Collect(maps.Values(store))
}
