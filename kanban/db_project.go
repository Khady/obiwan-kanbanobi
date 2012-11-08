package main

import (
	// "database/sql"
	// _ "github.com/bmizerany/pq"
)

type Project struct {
	Id         int
	Name       string
	admins_id  string
	Read       string
	Content    string
}
