package main

import (
	// "database/sql"
	// _ "github.com/bmizerany/pq"
)

type Script struct {
	Id         int
	script_type int
	object_id int
	column_name string
	Content    string
}
