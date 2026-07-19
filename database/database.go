package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connection()(*pgxpool.Pool, error) {
	url := "postgres://postgres:qwerty@localhost:5432/tododb"
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err.Error())
	}

	return conn, nil
}