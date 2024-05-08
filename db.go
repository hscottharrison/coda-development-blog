package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	API_URL = "https://yeeohvfasiodwextfhtj.supabase.co"
	API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InllZW9odmZhc2lvZHdleHRmaHRqIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTUwMDYxMzYsImV4cCI6MjAzMDU4MjEzNn0.93PnJeEsboDfyrXgQ8sBVdYCEW4xbd2Z8bfo2zjOo3k"
)

type DB interface {
	CreateCategory(*Category) error
	GetCategories() ([]*Category, error)
	CreatePost(*Post) error
	GetPosts() ([]*Post, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres.yeeohvfasiodwextfhtj password=MountLadyWashington2012 host=aws-0-us-west-1.pooler.supabase.com port=5432 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(0)

	return &PostgresStore{db: db}, nil
}
