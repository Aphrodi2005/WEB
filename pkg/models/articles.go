// article.go
package models

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var (
	ErrNoMovie   = errors.New("models: no matching movie found")
	ErrDuplicate = errors.New("models: duplicate movie title")
)

type Movie struct {
	ID          int
	Title       string
	Genre       string
	Rating      int
	SessionTime time.Time
}

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) Create(title, genre string, rating int, sessionTime time.Time) error {
	stmt := `INSERT INTO movies (title, genre, rating, sessionTime) VALUES (?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, title, genre, rating, sessionTime)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (m *MovieModel) Update(title, genre string, id, rating int, sessionTime time.Time) error {
	stmt := `UPDATE movies SET title=?, genre=?, rating=?, sessionTime=? WHERE id=?`
	_, err := m.DB.Exec(stmt, title, genre, rating, sessionTime, id)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}
func (m *MovieModel) Delete(id int) error {
	stmt := `DELETE FROM movies WHERE id=?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicate
		}
		return err
	}
	return nil
}
func isDuplicateError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Error 1062:")
}

func (m *MovieModel) Get(id int) (*Movie, error) {

	stmt := `SELECT id, title,  genre, rating,  sessionTime FROM movies WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	movie := &Movie{}
	err := row.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Rating, &movie.SessionTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoMovie
		}
		return nil, err
	}

	return movie, nil
}

func (m *MovieModel) Latest(int) ([]*Movie, error) {

	stmt := `SELECT id, title, genre, rating, sessionTime FROM movies ORDER BY sessionTime DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		movie := &Movie{}
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Rating, &movie.SessionTime)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
func (m *MovieModel) GetMovieByGenre(genre string) ([]*Movie, error) {
	query := `
        SELECT id, title, genre, rating, sessionTime
        FROM movies
        WHERE genre = ?
        ORDER BY sessionTime DESC
    `

	rows, err := m.DB.Query(query, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		movie := &Movie{}
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Rating, &movie.SessionTime)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
