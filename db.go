package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	// проверка соединения с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Println("Connected to PostgreSQL")
	return db, nil
}

func initDB(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id         SERIAL PRIMARY KEY,
			title      TEXT        NOT NULL,
			done       BOOLEAN     NOT NULL DEFAULT FALSE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);`
	_, err := db.Exec(query)
	return err
}
