package database

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type HTTPHandler struct {
	Pool *pgxpool.Pool
}

func StartDB() (*pgxpool.Pool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	//url := "postgres://kurswork:27121990@127.0.0.1/kurswork?sslmode=disable"
	url := os.Getenv("PG_DSN")
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 8
	cfg.MinConns = 4
	cfg.HealthCheckPeriod = 1 * time.Minute
	cfg.MaxConnLifetime = 24 * time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		Timeout:   cfg.ConnConfig.ConnectTimeout,
	}).DialContext

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	log.Print("db start")
	return dbpool, nil
}

func StopDB(ctx context.Context, pool *pgxpool.Pool) (err error) {

	for {
		select {
		case <-ctx.Done():
			pool.Close()
			log.Println("db stopped")
			return
		default:
			continue
		}
	}

}
