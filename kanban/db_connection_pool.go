package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

type InitFunction func(string) (*sql.DB, error)

type ConnectionPoolWrapper struct {
	size int
	conn chan *sql.DB
}

/**
Call the init function size times. If the init function fails during any call, then
the creation of the pool is considered a failure.
We call the same function size times to make sure each connection shares the same
state.
*/
func (p *ConnectionPoolWrapper) InitPool(size int, initfn InitFunction, url string) error {
	// Create a buffered channel allowing size senders
	p.conn = make(chan *sql.DB, size)
	for x := 0; x < size; x++ {
		conn, err := initfn(url)
		if err != nil {
			return err
		}
		var tmp int
		row := conn.QueryRow("select 1")
		err = row.Scan(&tmp)
		if err != nil {
			return err
		}
		// If the init function succeeded, add the connection to the channel
		p.conn <- conn
	}
	p.size = size
	return nil
}

func (p *ConnectionPoolWrapper) GetConnection() *sql.DB {
	return <-p.conn
}

func (p *ConnectionPoolWrapper) ReleaseConnection(conn *sql.DB) {
	p.conn <- conn
}

var dbPool = &ConnectionPoolWrapper{}
