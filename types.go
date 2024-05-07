package main

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstname"`
	LastLame  string    `json:"lastname"`
	CreatedAt time.Time `json:"createdat"`
}

type CreateCategoryRequest struct {
	CategoryName string `json:"categoryname"`
}

type Category struct {
	Id           int       `json:"id"`
	CategoryName string    `json:"categoryname"`
	CreatedAt    time.Time `json:"createdat"`
}

func NewCategory(CategoryName string) *Category {
	return &Category{
		CategoryName: CategoryName,
		CreatedAt:    time.Now().UTC(),
	}
}

type CreatePostRequest struct {
	PostTitle  string `json:"posttitle"`
	Content    string `json:"content"`
	CategoryId int    `json:"categoryid"`
	UserId     string `json:"userid"`
}

type Post struct {
	Id         int       `json:"id"`
	PostTitle  string    `json:"posttitle"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdat"`
	CategoryId int       `json:"categoryid"`
	UserId     string    `json:"userid"`
}

func NewPost(PostTitle string, Content string, CategoryId int, UserId string) *Post {
	return &Post{
		PostTitle:  PostTitle,
		Content:    Content,
		CategoryId: CategoryId,
		UserId:     UserId,
		CreatedAt:  time.Now().UTC(),
	}
}
