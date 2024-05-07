package main

import (
	"encoding/json"
	"net/http"
)

func (a *APIServer) HandleCreateCategory(w http.ResponseWriter, r *http.Request) error {
	createCategoryReq := new(CreateCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(createCategoryReq); err != nil {
		return err
	}

	category := NewCategory(createCategoryReq.CategoryName)
	a.store.CreateCategory(category)
	return WriteJSON(w, http.StatusOK, createCategoryReq)
}

func (a *APIServer) HandleGetCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := a.store.GetCategories()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, categories)
}
