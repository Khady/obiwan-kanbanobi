package main

type Column struct {
	Id         int
	Name       string
	Project_id uint32
	Content    string
	Tags       []string
	Scripts_id []uint32
	Write      []uint32
}
