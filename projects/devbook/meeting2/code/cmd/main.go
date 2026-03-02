package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"devbook_meeting2/code/pkg/api"
	"devbook_meeting2/code/pkg/devbook"
)

const ServerAddress = ":8080"

//go:embed static/index.html
var indexData []byte

func setupAndRunServer(store *devbook.Store, indexData []byte) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexData)
	})

	itemHandler := api.ItemHandler{Store: store}
	itemHandler.Register(mux)

	fmt.Println("Starting server on " + ServerAddress)
	return http.ListenAndServe(ServerAddress, mux)
}

func main() {
	// Setup starting data
	store := devbook.NewStore()
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

	if err := setupAndRunServer(store, indexData); err != nil {
		panic(err)
	}
}
