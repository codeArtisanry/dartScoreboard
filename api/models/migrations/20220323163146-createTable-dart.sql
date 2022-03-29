-- +migrate Up
CREATE TABLE IF NOT EXISTS dart(
    userId TEXT PRIMARY KEY,
    userEmail TEXT,
    userPicture TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS dart();
