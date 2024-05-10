package main

import (
	"fmt"
)

func (s *PostgresStore) CreatePost(post *Post) error {
	query := `INSERT INTO posts
  (posttitle, content, categoryid, userid, createdat, imageurl, isdraft)
  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, queryErr := s.db.Exec(query, post.PostTitle, post.Content, post.CategoryId, post.UserId, post.CreatedAt, post.ImageUrl, post.IsDraft)

	if queryErr != nil {
		fmt.Println("Error in creating post: ", queryErr)
		return queryErr
	}

	return nil
}

func (s *PostgresStore) GetPosts(getDrafts bool) ([]*Post, error) {
	query := `SELECT * FROM posts ORDER BY createdat DESC`

	if !getDrafts {
		query = `SELECT * FROM posts WHERE isdraft=false ORDER BY createdat DESC`
	}

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
			&post.ImageUrl,
			&post.IsDraft)

		if err != nil {
			fmt.Println("Error in scanning rows: ", err)
			return nil, err
		}
		posts = append(posts, post)

	}

	return posts, nil
}

func (s *PostgresStore) UpdatePost(id string, post *Post) error {
	query := `UPDATE posts SET posttitle=$1, content=$2, categoryid=$3, userid=$4, imageurl=$5, isdraft=$6 WHERE id=$7`

	_, queryErr := s.db.Exec(query, post.PostTitle, post.Content, post.CategoryId, post.UserId, post.ImageUrl, post.IsDraft, id)

	if queryErr != nil {
		fmt.Println("Error in updating post: ", queryErr)
		return queryErr
	}

	return nil
}

func (s *PostgresStore) DeletePost(id string) error {
	query := `DELETE FROM posts WHERE id=$1`

	_, queryErr := s.db.Exec(query, id)

	if queryErr != nil {
		fmt.Println("Error in deleting post: ", queryErr)
		return queryErr
	}

	return nil
}
