package database

import (
	"database/sql"
	"fmt"

	"github.com/Virean196/bos-pt-eng-translator/internal/translator"
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

func (s *DB) CreatePhrase(phrase string, langPair string) (string, error) {
	translation, err := translator.GetTranslation(phrase, langPair)
	if err != nil {
		return "", err
	}
	_, err = s.db.Exec(
		"INSERT INTO phrases (input, translation, langPair) VALUES (?, ?, ?)",
		phrase, translation, langPair,
	)
	if err != nil {
		return "", fmt.Errorf("Could not execute query: %w", err)
	}
	return translation, nil
}
