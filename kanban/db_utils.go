package main

import (
	"database/sql"
	"github.com/bmizerany/pq"
)

func db_open(url string) (*sql.DB, error) {
	conn_str, err := pq.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", conn_str);
}













