package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	ctx := context.Background()

	// Формирование строки подключения
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, username, password, dbname,
	)

	// Парсинг конфигурации пула подключений
	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse pool config: %v\n", err)
		os.Exit(1)
	}

	// Подключение к базе данных
	pool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close() // Важно закрыть пул подключений при завершении программы

	fmt.Println("Success connection")
}
