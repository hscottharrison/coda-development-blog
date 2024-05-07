package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *APIServer) HandleGetPosts(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "return all posts")
	return nil
}

func (a *APIServer) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("create post")
	createPostReq := new(CreatePostRequest)
	fmt.Println("createPostReq", createPostReq)
	if err := json.NewDecoder(r.Body).Decode(createPostReq); err != nil {
		return err
	}
	post := NewPost(createPostReq.PostTitle, createPostReq.Content, createPostReq.CategoryId, createPostReq.UserId)
	a.store.CreatePost(post)
	return WriteJSON(w, http.StatusOK, createPostReq)
}
