package main

import "time"

type CreateCategoryRequest struct {
	CategoryName string `json:"categoryName"`
}

type Category struct {
	Id           int       `json:"id"`
	CategoryName string    `json:"categoryName"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewCategory(CategoryName string) *Category {
	return &Category{
		CategoryName: CategoryName,
		CreatedAt:    time.Now().UTC(),
	}
}

type CreatePostRequest struct {
	PostTitle  string `json:"postTitle"`
	Content    string `json:"content"`
	CategoryId int    `json:"categoryId"`
	UserId     string `json:"userId"`
	ImageUrl   string `json:"imageUrl"`
	IsDraft    bool   `json:"isDraft"`
}

type Post struct {
	Id         int       `json:"id"`
	PostTitle  string    `json:"postTitle"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	CategoryId int       `json:"categoryId"`
	UserId     string    `json:"userId"`
	ImageUrl   string    `json:"imageUrl"`
	IsDraft    bool      `json:"isDraft"`
}

func NewPost(p *CreatePostRequest) *Post {
	return &Post{
		PostTitle:  p.PostTitle,
		Content:    p.Content,
		CategoryId: p.CategoryId,
		UserId:     p.UserId,
		ImageUrl:   p.ImageUrl,
		CreatedAt:  time.Now().UTC(),
		IsDraft:    p.IsDraft,
	}
}
