-- +migrate Up
CREATE TABLE IF NOT EXISTS dart(
    userId TEXT PRIMARY KEY,
    userEmail TEXT,
    userPicture TEXT
); CREATE TABLE IF NOT EXISTS game(
    GameName TEXT PRIMARY KEY,
    GameType TEXT,
    PlayersNames TEXT,
    GameTargetScore TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS dart(); DROP TABLE IF EXISTS game();
