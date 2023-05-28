package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres() (*sqlx.DB, error) {
	// создаем объект для работы с бд
	// поменять драйвер на mysql
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user= dbname=postgres password= sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}
	// проверяем, что доходят запросы
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}
	return db, nil
}
