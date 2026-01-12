package db

import (
	"database/sql"
	"fmt"
	"net/url"
)

func NewMySQLFromURI(uri string) (*sql.DB, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	user := u.User.Username()
	password, _ := u.User.Password()
	host := u.Host
	dbName := u.Path[1:] // enlève le /

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		user,
		password,
		host,
		dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
