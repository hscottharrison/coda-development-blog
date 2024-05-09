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

func (a *APIServer) HandleUpdatePost(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	post := new(Post)
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		return err
	}

	err := a.store.UpdatePost(id, post)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, post)
}

func (a *APIServer) HandleDeletePost(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	err := a.store.DeletePost(id)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, "Post deleted successfully")
}
