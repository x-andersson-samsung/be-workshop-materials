package http

import (
	"net/http"

	"devbook_meeting2/code/pkg/devbook"
)

type ItemHandler struct {
	Store devbook.Store
}

func (i *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (i *ItemHandler) Remove(w http.ResponseWriter, r *http.Request) {

}

func (i *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (i *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
