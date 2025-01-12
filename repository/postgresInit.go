package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	instance *sql.DB
	once     sync.Once
)

func PostgresDbInit() (*sql.DB, error) {
	var err error
	once.Do(func() {
		psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

		instance, err = sql.Open("postgres", psqlconn)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
			return
		}

		if err = instance.Ping(); err != nil {
			_ = instance.Close() //Close db if ping failed
			log.Fatalf("Failed to ping database: %v", err)
			return
		}

	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}
