
-- +migrate Up
CREATE TABLE IF NOT EXISTS dart(
    id INTEGER PRIMARY KEY,
    token INTEGER,
    player_name TEXT,
    player_email TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS dart();
