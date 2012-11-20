package main

import (
	"database/sql"
	"github.com/bmizerany/pq"
	"strings"
)

func db_open(url string) (*sql.DB, error) {
	conn_str, err := pq.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", conn_str)
}

func getIntSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, id int) ([]int, error) {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var cell string
	var cell_ids []int
	row := db.QueryRow("select $1 from $2 where id = $3", column_name, base_name, id)
	if err := row.Scan(&cell); err != nil {
		return cell_ids, err
	}
	cell_ids = SInt_of_SString(strings.Split(cell, " "))
	return cell_ids, nil
}

func updateIntSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, cell []int, id int) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	new_cell := strings.Join(SString_of_SInt(cell), " ")
	_, err := db.Exec("update $1 set $2 = $3 where id = $4", column_name, base_name, new_cell, id)
	return err
}

func getStringSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, id int) ([]string, error) {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var cell string
	var cell_ids []string
	row := db.QueryRow("select $1 from $2 where id = $3", column_name, base_name, id)
	if err := row.Scan(&cell); err != nil {
		return cell_ids, err
	}
	cell_ids = strings.Split(cell, " ")
	return cell_ids, nil
}

func updateStringSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, cell []string, id int) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	new_cell := strings.Join(cell, " ")
	_, err := db.Exec("update $1 set $2 = $3 where id = $4", column_name, base_name, new_cell, id)
	return err
}