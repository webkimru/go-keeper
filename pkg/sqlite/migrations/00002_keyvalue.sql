-- +goose Up
CREATE TABLE IF NOT EXISTS keyvalues (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TEXT,
    updated_at TEXT
);

-- +goose Down
DROP TABLE keyvalues;