-- +goose Up
CREATE TABLE phrases (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  input text NOT NULL,
  translation text NOT NULL,
  langPair text NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE phrases;
