package main

import (
	"database/sql"
	"github.com/bmizerany/pq"
	"sort"
	"strings"
)

func db_open(url string) (*sql.DB, error) {
	conn_str, err := pq.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", conn_str)
}

func getUInt32SliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, id uint32) ([]uint32, error) {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var cell string
	var cell_ids []uint32
	row := db.QueryRow("select $1 from $2 where id = $3", column_name, base_name, id)
	if err := row.Scan(&cell); err != nil {
		return cell_ids, err
	}
	cell_ids = SUInt32_of_SString(strings.Split(cell, ","))
	return cell_ids, nil
}

func updateUInt32SliceCell(p *ConnectionPoolWrapper, column_name string, base_name string, cell []uint32, id uint32) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	new_cell := strings.Join(SString_of_SUInt32(cell), ",")
	_, err := db.Exec("update $1 set $2 = $3 where id = $4", base_name, column_name, new_cell, id)
	return err
}

func getStringSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, id uint32) ([]string, error) {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var cell string
	var cell_ids []string
	row := db.QueryRow("select $1 from $2 where id = $3", column_name, base_name, id)
	if err := row.Scan(&cell); err != nil {
		return cell_ids, err
	}
	cell_ids = strings.Split(cell, ",")
	return cell_ids, nil
}

func updateStringSliceCell(p *ConnectionPoolWrapper, base_name string, column_name string, cell []string, id uint32) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	new_cell := strings.Join(cell, ",")
	_, err := db.Exec("update $1 set $2 = $3 where id = $4", column_name, base_name, new_cell, id)
	return err
}

func addIdInCell(p *ConnectionPoolWrapper, id uint32, elem_id uint32, table string, column string) error {
	cell_ids, err := getUInt32SliceCell(p, table, column, elem_id)
	if err != nil {
		return err
	}
	id_idx := sort.Search(len(cell_ids), func(i int) bool { return cell_ids[i] == id }) // C'est pas un peu crade ce search niveau complexite ?
	if id_idx == len(cell_ids) {
		cell_ids = append(cell_ids, id)
		err = updateUInt32SliceCell(p, table, column, cell_ids, elem_id)
	}
	return err
}

func delIdInCell(p *ConnectionPoolWrapper, id uint32, elem_id uint32, table string, column string) error {
	cell_ids, err := getUInt32SliceCell(p, table, column, elem_id)
	if err != nil {
		return err
	}
	id_idx := sort.Search(len(cell_ids), func(i int) bool { return cell_ids[i] == id }) // C'est pas un peu crade ce search niveau complexite ?
	if id_idx != len(cell_ids) {
		new_ids := []uint32{}
		new_ids = append(new_ids, cell_ids[:id_idx]...)
		new_ids = append(new_ids, cell_ids[id_idx+1:]...)
		return updateUInt32SliceCell(p, table, column, new_ids, elem_id)
	}
	return err
}
