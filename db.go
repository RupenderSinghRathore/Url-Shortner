package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func UpdateDb(url, urlCode string) error {
	db, err := sql.Open("sqlite3", "./url.sql")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				url TEXT UNIQUE,
				shorturl TEXT UNIQUE
	)`)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT OR IGNORE INTO users(url, shorturl) VALUES (?, ?)", url, urlCode)
	if err != nil {
		return err
	}

	return nil
}

func RetriveDb(code string) (string, error) {
	db, err := sql.Open("sqlite3", "./url.sql")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var url string
	err = db.QueryRow(`SELECT url FROM users WHERE shorturl = ?`, code).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
