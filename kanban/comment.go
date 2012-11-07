package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

type Comment struct {
	Id        int
	Content   string
	Cards_id  int
	Author_id int
}

func UpdateComment(db *sql.DB, c *Comment) error {
	_, err := db.Exec(`update cards set content = $1, cards_id = $2, author_id = $3, where id = $4`,
		c.Content, c.Cards_id, c.Author_id, c.Id)
	return err
}

func AddComment(db *sql.DB, c *Comment) error {
	_, err := db.Exec("INSERT INTO users(content, cards_id, author_id) VALUES($1, $2, $3);",
		c.Content, c.Cards_id, c.Author_id)
	return err
}

func DelComment(db *sql.DB, id int) error {
	_, err := db.Exec("delete from comments where id = $1", id)
	return err
}
