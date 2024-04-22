package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPgxConn pool
func NewPgxConn(env *Env) (*pgxpool.Pool, error) {
	username := env.DBUser
	password := env.DBPass
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

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

	return pool , nil
}

func Close(p *pgxpool.Pool) {
	if p != nil {
		p.Close()
	}
}
