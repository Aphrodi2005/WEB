// category.go
package models

import (
	"database/sql"
	"errors"
)

var ErrNoCategory = errors.New("models: no matching category found")

type Category struct {
	ID   int
	Name string
	// Add other fields as needed
}

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) Get(id int) (*Category, error) {
	// Implement the logic to retrieve a single category by ID from the database
	stmt := `SELECT id, name FROM categories WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	category := &Category{}
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoCategory
		}
		return nil, err
	}

	return category, nil
}

func (m *CategoryModel) All() ([]*Category, error) {
	// Implement the logic to retrieve all categories from the database
	stmt := `SELECT id, name FROM categories`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*Category

	for rows.Next() {
		category := &Category{}
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (m *CategoryModel) Articles(categoryID int) ([]*Article, error) {
	// Implement the logic to retrieve articles for a specific category from the database
	stmt := `SELECT id, title, content, created FROM articles WHERE category_id = ?`
	rows, err := m.DB.Query(stmt, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Created)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
