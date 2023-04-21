package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var lobbyTable = `
CREATE TABLE IF NOT EXISTS lobbies(
	id TEXT PRIMARY KEY,
	category_id TEXT, 			/* immutable */
	guild_id TEXT, 				/* immutable */
	template TEXT,				/* mutable, default NULL */
	capacity INTEGER			/* mutable, default NULL */
);`

var channelTable = `
CREATE TABLE IF NOT EXISTS channels(
	id TEXT PRIMARY KEY,
	parent_id TEXT NOT NULL		/* immutable */
);`

var channelMembersTable = `
CREATE TABLE IF NOT EXISTS channel_members(
	user_id TEXT PRIMARY KEY,
	channel_id TEXT NOT NULL,
	guild_id TEXT NOT NULL
);`

func Load() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return nil, fmt.Errorf("open sql: %w", err)
	}

	_, err = db.Exec(lobbyTable)
	if err != nil {
		return nil, fmt.Errorf("create lobby table: %w", err)
	}

	_, err = db.Exec(channelTable)
	if err != nil {
		return nil, fmt.Errorf("create channel table: %w", err)
	}

	_, err = db.Exec(channelMembersTable)
	if err != nil {
		return nil, fmt.Errorf("create channel members table: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("verify db connection: %w", err)
	}

	log.Println("Storage loaded!")
	return db, nil
}
