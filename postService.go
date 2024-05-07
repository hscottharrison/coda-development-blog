package main

import "log"

func (s *PostgresStore) CreatePost(post *Post) error {
	query := `INSERT INTO posts
  (posttitle, content, categoryid, userid, createdat)
  VALUES ($1, $2, $3, $4, $5)`

	rows, queryErr := s.db.Exec(query, post.PostTitle, post.Content, post.CategoryId, post.UserId, post.CreatedAt)

	if queryErr != nil {
		log.Fatal("ERROR in executing query: ", queryErr)
		return queryErr
	}

	response := rows // Assuming you want to assign the rows to the response variable

	// Rest of the code...
	log.Println(response)

	return nil
}
