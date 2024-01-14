// article.go
package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

var (
	ErrNoArticle = errors.New("models: no matching article found")
	ErrDuplicate = errors.New("models: duplicate article title")
)

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type ArticleModel struct {
	DB *sql.DB
}

// ... (existing methods)

func (m *ArticleModel) Create(title, content string) error {
	stmt := `INSERT INTO articles (title, content, created) VALUES (?, ?, UTC_TIMESTAMP())`
	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (m *ArticleModel) Update(id int, title, content string) error {
	stmt := `UPDATE articles SET title=?, content=?, created=UTC_TIMESTAMP() WHERE id=?`
	_, err := m.DB.Exec(stmt, title, content, id)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (m *ArticleModel) Delete(id int) error {
	stmt := `DELETE FROM articles WHERE id=?`
	_, err := m.DB.Exec(stmt, id)
	return err
}

// Helper function to check if a MySQL error is a duplicate entry error
func isDuplicateError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Error 1062:")
}

func (m *ArticleModel) Get(id int) (*Article, error) {
	// Implement the logic to retrieve an article by ID from the database
	stmt := `SELECT id, title, content, created FROM articles WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	article := &Article{}
	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoArticle
		}
		return nil, err
	}

	return article, nil
}

func (m *ArticleModel) Latest() ([]*Article, error) {
	// Implement the logic to retrieve the latest articles from the database
	stmt := `SELECT id, title, content, created FROM articles ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
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
