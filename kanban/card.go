package main

type Card struct {
	Id         int
	Name       string
	Content    string
	Column_id  int
	Project_id int
	Tags       []string
	User_id    []int
	Scripts_id []int
	Write      []int
}
