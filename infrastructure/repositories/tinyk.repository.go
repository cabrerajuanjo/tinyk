package repository

import (
	"database/sql"
)

func CreateUrl(db *sql.DB, url string, smallUrl string) error {
	_, err := db.Exec("INSERT INTO url (url, smallUrl) VALUES ($1, $2)", url, smallUrl)
	return err
}
