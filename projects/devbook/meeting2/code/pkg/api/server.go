package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"devbook_meeting2/code/pkg/devbook"
)

type ItemHandler struct {
	Store *devbook.Store
}

func (i *ItemHandler) Register(r *http.ServeMux) {
	r.HandleFunc("POST /items", i.Create)
	r.HandleFunc("GET /items", i.List)
	r.HandleFunc("PUT /items/{id}", i.Update)
	r.HandleFunc("DELETE /items/{id}", i.Delete)
}

func (i *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item devbook.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	i.Store.Add(item)

	w.WriteHeader(http.StatusNoContent)
}

func (i *ItemHandler) parsePagination(r *http.Request) (limit int, offset int, err error) {
	if err := r.ParseForm(); err != nil {
		return 0, 0, err
	}

	limitStr := r.Form.Get("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, 0, err
		}
	}

	offsetStr := r.Form.Get("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return 0, 0, err
		}
	}

	return limit, offset, nil
}

func (i *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	paginationLimit, paginationOffset, err := i.parsePagination(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items := i.Store.List()

	// Apply offset
	if paginationOffset > 0 {
		items = items[min(len(items), paginationOffset):]
	}

	// Apply limit
	if paginationLimit > 0 {
		items = items[:min(len(items), paginationLimit)]
	}

	payload, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (i *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (i *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	i.Store.DeleteByID(id)
}
