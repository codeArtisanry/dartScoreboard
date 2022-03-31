-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    avatar TEXT,
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
    creater_email TEXT,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    )
);
CREATE TABLE IF NOT EXISTS scores(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT
    round INTEGER,
    throw INTEGER,
    point INTEGER,
    game_id INTEGER,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(game_id) REFERENCES games (id)
    FOREIGN KEY(user_id) REFERENCES games (id)
);
CREATE TABLE IF NOT EXISTS players(
    user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    updated_at DATETIME DEFAULT (
        STRFTIME('%d-%m-%Y   %H:%M:%S', 'NOW', 'localtime')
    ),
    FOREIGN KEY(user_id) REFERENCES users (id),
    PRIMARY KEY(email, game_id)
);
-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS scores;
DROP TABLE IF EXISTS players;
