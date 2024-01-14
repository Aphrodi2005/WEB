// contact.go
package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoContact = errors.New("models: no matching contact found")

type Contact struct {
	ID        int
	Name      string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ContactModel struct {
	DB *sql.DB
}

func (m *ContactModel) Get(id int) (*Contact, error) {
	// Implement the logic to retrieve a single contact by ID from the database
	stmt := `SELECT id, name, location, created_at, updated_at FROM contacts WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	contact := &Contact{}
	err := row.Scan(&contact.ID, &contact.Name, &contact.Location, &contact.CreatedAt, &contact.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoContact
		}
		return nil, err
	}

	return contact, nil
}

func (m *ContactModel) All() ([]*Contact, error) {
	// Implement the logic to retrieve all contacts from the database
	stmt := `SELECT id, name, location, created_at, updated_at FROM contacts`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*Contact

	for rows.Next() {
		contact := &Contact{}
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Location, &contact.CreatedAt, &contact.UpdatedAt)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}
