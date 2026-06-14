package database

import (
	"database/sql"
	"fmt"
)

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{db: db}
}

func (s *DB) GetPhrase(input string) (string, error) {
	var translation string
	err := s.db.QueryRow("SELECT translation FROM phrases WHERE input = ?", input).Scan(&translation)
	if err == sql.ErrNoRows {
		return "", sql.ErrNoRows
	}
	if err != nil {
		return "", fmt.Errorf("Could not get phrase: %w", err)
	}
	return translation, nil
}

func (s *DB) CreatePhrase(phrase string) error {
	var translation = fmt.Sprintf("%s.translated", phrase)
	_, err := s.db.Exec(
		"INSERT INTO phrases (input, translation) VALUES (?, ?)",
		phrase, translation,
	)
	if err != nil {
		return fmt.Errorf("Could not execute query: %w", err)
	}
	return nil
}
