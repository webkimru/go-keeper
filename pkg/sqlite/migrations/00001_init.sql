-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    status NOT NULL,
    created_at TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS user_idx ON users (login);

-- +goose Down
DROP TABLE users;