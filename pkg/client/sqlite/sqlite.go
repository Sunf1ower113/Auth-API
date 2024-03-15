package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewDB(driver, name string) (db *sql.DB, err error) {
	db, err = sql.Open(driver, name)
	if err != nil {
		log.Println("error create db")
		return
	}
	log.Println("create db: OK")
	err = createTable(db)
	if err != nil {
		log.Println("error create tables")
		return
	}
	return
}
func createTable(db *sql.DB) error {
	var query []string
	users := `
CREATE TABLE IF NOT EXISTS users(
user_id INTEGER PRIMARY KEY AUTOINCREMENT,
username TEXT DEFAULT '',
email TEXT UNIQUE NOT NULL,
password TEXT NOT NULL,
phone_number TEXT DEFAULT '',
birth_date DATE DEFAULT '',
role TEXT CHECK (role IN ('admin', 'user')) DEFAULT 'user'
)
`
	movies := `
CREATE TABLE IF NOT EXISTS movies(
movie_id INTEGER PRIMARY KEY AUTOINCREMENT,
title TEXT,
category TEXT CHECK (category IN ('Action', 'Animation', 'Comedy', 'Crime', 'Drama', 'Experimental', 'Fantasy', 'Historical', 'Horror', 'Romance', 'Science Fiction', 'Thriller', 'Western', 'Other')),
type TEXT CHECK (type IN ('Movies', 'TV Shows', 'Videos')),
age_group TEXT CHECK (age_group IN ('G', 'PG', 'PG-13','R', 'NC-17')),
year INTEGER,
timing INTEGER,
tags TEXT,
description TEXT,
director TEXT,
producer TEXT
)
`
	query = append(query, users, movies)
	for _, v := range query {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func migrate(db *sql.DB) error {
	var query []string
	users := ``
	query = append(query, users)
	for _, v := range query {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}
