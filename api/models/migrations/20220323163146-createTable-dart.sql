-- +migrate Up
CREATE TABLE IF NOT EXISTS dart(
    userId TEXT PRIMARY KEY,
    userEmail TEXT,
    userPicture TEXT
); CREATE TABLE IF NOT EXISTS game(
    gameName TEXT PRIMARY KEY,
    gameType TEXT,
    playersNames TEXT,
    gameTargetScore TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS dart(); DROP TABLE IF EXISTS game();
