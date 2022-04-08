-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    avatar_url TEXT,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    )
);
CREATE TABLE IF NOT EXISTS games(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    type TEXT,
    status TEXT,
    creater_user_id TEXT,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(creater_user_id) REFERENCES users (id)
);
CREATE TABLE IF NOT EXISTS rounds(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    round INTEGER,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(game_player_id) REFERENCES game_players (id)
);
CREATE TABLE IF NOT EXISTS scores(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    round_id INTEGER,
    player_id INTEGER,
    throw INTEGER,
    score INTEGER,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(round_id) REFERENCES rounds (id)
);
CREATE TABLE IF NOT EXISTS game_players(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    game_id INTEGER,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(user_id) REFERENCES users (id),
    FOREIGN KEY(game_id) REFERENCES games (id)
);
-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS rounds;
DROP TABLE IF EXISTS scores;
DROP TABLE IF EXISTS game_players;
