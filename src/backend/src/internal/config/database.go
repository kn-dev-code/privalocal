package config

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func PostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	_ = godotenv.Load()

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return nil, fmt.Errorf("Invalid database url")
	}

	pool, err := pgxpool.New(ctx, databaseUrl)

	if err != nil {
		return nil, fmt.Errorf("created pgx pool error: %w\n", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w\n", err)
	}

	return pool, nil
}