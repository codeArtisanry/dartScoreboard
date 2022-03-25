
-- +migrate Up
CREATE TABLE IF NOT EXISTS dart(
    id INTEGER PRIMARY KEY,
    token TEXT,
    player_name TEXT,
    player_email TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS dart();
