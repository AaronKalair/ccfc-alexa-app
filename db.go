package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Match struct {
	match string
	date  time.Time
}

func InitDb(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func ReadItem(db *sql.DB) []Match {
	sql_readall := `
	SELECT match, date FROM matches
	WHERE date >= DATETIME('now') ORDER BY datetime(date) ASC LIMIT 1
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Match
	for rows.Next() {
		item := Match{}
		err2 := rows.Scan(&item.match, &item.date)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS matches(
		id INTEGER PRIMARY KEY,
		match TEXT,
		date DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func StoreItem(db *sql.DB, match Match) {
	sql_additem := `
	INSERT OR REPLACE INTO matches(
		match,
		date
	) values(?, ?)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(match.match, match.date)
	if err2 != nil {
		panic(err2)
	}
}
