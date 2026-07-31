package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connection()(*pgxpool.Pool, error) {
	envErr := godotenv.Load()
	if envErr != nil { 
		log.Fatalf("env not found")
	}
	var err error
	url := os.Getenv("DB_URL")
	DB, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Bad connection to database %v", err.Error())
	}

	return DB, nil
}