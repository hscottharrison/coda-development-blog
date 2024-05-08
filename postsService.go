package main

import (
	"fmt"
)

func (s *PostgresStore) CreatePost(post *Post) error {
	query := `INSERT INTO posts
  (posttitle, content, categoryid, userid, createdat, imageurl)
  VALUES ($1, $2, $3, $4, $5, $6)`

	_, queryErr := s.db.Exec(query, post.PostTitle, post.Content, post.CategoryId, post.UserId, post.CreatedAt, post.ImageUrl)

	if queryErr != nil {
		fmt.Println("Error in creating post: ", queryErr)
		return queryErr
	}

	return nil
}

func (s *PostgresStore) GetPosts() ([]*Post, error) {
	query := `SELECT * FROM posts ORDER BY createdat DESC`

	rows, queryErr := s.db.Query(query)

	if queryErr != nil {
		return nil, queryErr
	}

	posts := []*Post{}
	for rows.Next() {
		post := new(Post)
		err := rows.Scan(
			&post.Id,
			&post.CreatedAt,
			&post.PostTitle,
			&post.Content,
			&post.CategoryId,
			&post.UserId,
			&post.ImageUrl)

		if err != nil {
			fmt.Println("Error in scanning rows: ", err)
			return nil, err
		}
		posts = append(posts, post)

	}

	return posts, nil
}
