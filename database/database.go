package database

import (
	"context"
	"log"
	"net"
	"time"
        "os"
	"github.com/jackc/pgx/v4/pgxpool"
	"short/logrus"
)

type HTTPHandler struct {
	Pool *pgxpool.Pool
}

func StartDB() (*pgxpool.Pool, error) {
	
	logFile, hlog := logrus.LogInit()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	url := os.Getenv("DATABASE_URL")
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		hlog.Warnf("configuration pgxpool not valid (func StartDB, package database), %s",err)
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
	
	err = logFile.Close()
	if err != nil {
		hlog.Errorf("Файл логов не закрылся %s", err)
	}
	
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
