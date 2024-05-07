package main

import "log"

func (s *PostgresStore) CreateCategory(c *Category) error {
	query := ` INSERT INTO categories 
	(categoryname, createdat)
	VALUES ($1, $2)`

	rows, queryErr := s.db.Exec(query, c.CategoryName, c.CreatedAt)

	if queryErr != nil {
		log.Fatal("ERROR in executing query: ", queryErr)
		return queryErr
	}

	response := rows // Assuming you want to assign the rows to the response variable

	// Rest of the code...
	log.Println(response)

	return nil
}

func (s *PostgresStore) GetCategories() ([]*Category, error) {
	query := `SELECT * FROM categories`

	rows, queryErr := s.db.Query(query)

	if queryErr != nil {
		log.Fatal("ERROR in executing query: ", queryErr)
		return nil, queryErr
	}

	categories := []*Category{}
	for rows.Next() {
		category := new(Category)
		err := rows.Scan(
			&category.Id,
			&category.CreatedAt,
			&category.CategoryName)

		if err != nil {
			return nil, err
		}
		categories = append(categories, category) // Fix: Use the correct syntax for appending to a slice
	}

	return categories, nil
}
