package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB interface {
	CreateCategory(*Category) error
	GetCategories() ([]*Category, error)
	CreatePost(*Post) error
	GetPosts() ([]*Post, error)
	UpdatePost(string, *Post) error
	DeletePost(string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	godotenv.Load()
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	connStr := fmt.Sprintf("user=%s password=%s host=aws-0-us-west-1.pooler.supabase.com port=5432 dbname=postgres sslmode=disable", postgresUser, postgresPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(0)

	return &PostgresStore{db: db}, nil
}
