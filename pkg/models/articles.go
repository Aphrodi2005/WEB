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
	ID       int
	Title    string
	Content  string
	Created  time.Time
	Category sql.NullString // Используем sql.NullString для учета NULL
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) Create(title, content, category string) error {
	stmt := `INSERT INTO articles (title, content, category, created) VALUES (?, ?, ?, UTC_TIMESTAMP())`

	_, err := m.DB.Exec(stmt, title, content, category)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (m *ArticleModel) Update(title, content, category string, id int) error {
	stmt := `UPDATE articles SET title=?, content=?, category=?, created=UTC_TIMESTAMP() WHERE id=?`
	_, err := m.DB.Exec(stmt, title, content, category, id)
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
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

// Helper function to check if a MySQL error is a duplicate entry error
func isDuplicateError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Error 1062:")
}

func (m *ArticleModel) Get(id int) (*Article, error) {

	stmt := `SELECT id, title,  content, category,  created FROM articles WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	article := &Article{}
	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoArticle
		}
		return nil, err
	}

	return article, nil
}

func (m *ArticleModel) Latest(int) ([]*Article, error) {

	stmt := `SELECT id, title, content, category, created FROM articles ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Created)
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
func (m *ArticleModel) GetArticlesByCategory(category string) ([]*Article, error) {
	query := `
        SELECT id, title, content, category
        FROM articles
        WHERE category = ?
        ORDER BY created DESC
    `

	rows, err := m.DB.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Category)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
