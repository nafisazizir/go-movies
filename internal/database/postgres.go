package database

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(user, password, dbname string) (*Postgres, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}
