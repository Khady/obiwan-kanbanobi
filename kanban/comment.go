package main

import "time"

type Comment struct {
	Id           int
	Content      string
	Cards_id     int
	Author_id    int
	Comment_date time.Time
}
