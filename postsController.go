package main

import (
	"encoding/json"
	"net/http"
)

func (a *APIServer) HandleGetPosts(w http.ResponseWriter, r *http.Request) error {
	posts, err := a.store.GetPosts()

	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, posts)
}

func (a *APIServer) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	createPostReq := new(CreatePostRequest)
	if err := json.NewDecoder(r.Body).Decode(createPostReq); err != nil {
		return err
	}
	post := NewPost(createPostReq)
	a.store.CreatePost(post)
	return WriteJSON(w, http.StatusOK, createPostReq)
}
