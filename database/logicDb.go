package database

import (
	"context"
	"fmt"
	"log"
	"short/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertData(ctx context.Context, pool *pgxpool.Pool, userURL string, newURL string) (err error) {
	const sql = `
insert into form (userurl, newurl) values
($1, $2);`

	addData := &models.Form{UserURL: userURL, NewURL: newURL}

	_, err = pool.Exec(ctx, sql, addData.UserURL, addData.NewURL)
	if err != nil {
		return fmt.Errorf("error database exec, %s", err)
	}

	return nil
}

func InsertStat(ctx context.Context, pool *pgxpool.Pool, ip string, time time.Time, newurl string) (err error) {

	const sql = `
insert into statistic (ip, time,newurl_form) values
($1, $2,$3);`

	_, err = pool.Exec(ctx, sql, ip, time, newurl)
	if err != nil {
		return fmt.Errorf("error database exec, %s", err)

	}

	return nil
}

func ShowAllResult(ctx context.Context, pool *pgxpool.Pool) ([]models.Form, error) {

	rows, err := pool.Query(ctx, "select  id, userurl, newurl from form;")
	if err != nil {
		log.Println("not show all results")
		rows.Close()
		return nil, err
	}
	defer rows.Close()

	allResultsSlice := []models.Form{}

	for rows.Next() {
		f := models.Form{}
		err = rows.Scan(&f.ID, &f.UserURL, &f.NewURL)
		if err != nil {
			log.Println(err)
			rows.Close()
			return nil, err
		}
		allResultsSlice = append(allResultsSlice, f)
	}
	return allResultsSlice, nil
}

func SearchRowStat(ctx context.Context, dbpool *pgxpool.Pool, name string) ([]models.Stat, error) {
	const sql = `
select
ip,
time
from statistic
where newurl_form like $1
order by newurl_form asc;
`
	pattern := "http://127.1.1.0:2000/s/" + name
	rows, err := dbpool.Query(ctx, sql, pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()
	var hints []models.Stat
	for rows.Next() {
		var hint models.Stat
		err = rows.Scan(&hint.IP, &hint.Time)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		hints = append(hints, hint)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}
	return hints, nil
}

func Search(ctx context.Context, pool *pgxpool.Pool, name string) (models.Form, error) {
	const sql = `
select userurl, newurl from form where newurl like $1 order by userurl;`
	pattern := "http://127.1.1.0:2000/s/" + name

	rows, err := pool.Query(ctx, sql, pattern)
	if err != nil {
		return models.Form{}, fmt.Errorf("failed to query data: %w", err)
	}

	defer rows.Close()

	var hint models.Form
	for rows.Next() {

		err = rows.Scan(&hint.UserURL, &hint.NewURL)
		if err != nil {
			return models.Form{}, fmt.Errorf("failed to scan row: %w", err)
		}

	}

	if rows.Err() != nil {
		return models.Form{}, fmt.Errorf("failed to read response: %w", rows.Err())
	}

	return hint, nil
}
