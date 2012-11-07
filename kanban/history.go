package main

import (
	// "database/sql"
	// _ "github.com/bmizerany/pq"
)

type History struct {
	Id         int
	change_type int
	object_id int
	column_name string
	Content    string
}
