package devbook

type Store map[string]Item

func (store Store) GetByName(name string) (Item, bool) {
	item, ok := store[name]
	return item, ok
}

func (store Store) Add(item Item) {
	store[item.Name] = item
}

func (store Store) DeleteByName(name string) {
	delete(store, name)
}

func (store Store) List() []Item {
	// Expressive way
	arr := make([]Item, 0, len(store))
	for _, item := range store {
		arr = append(arr, item)
	}

	return arr

	// Compact way
	//return slices.Collect(maps.Values(store))
}
