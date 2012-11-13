package main

import (
	"testing"
)

func Test_UpdateComment(t *testing.T) {
	if err := dbPool.InitPool(2, db_open, INFO_CONNECT); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	db.Exec(`INSERT INTO comments(content, cards_id, author_id)
 VALUES('super content', '10', '10');`)
	row := db.QueryRow("select id from comments where content = 'super content'")
	var id int
	row.Scan(&id)
	oldc := &Comment{id, "new super content", 5, 5}
	if err := oldc.Update(dbPool); err != nil {
		t.Error(err)
	}
	row = db.QueryRow("select * from comments where id = $1", oldc.Id)
	c := &Comment{}
	row.Scan(&c.Id, &c.Content, &c.Cards_id, &c.Author_id)
	if c.Content != "new super content" || c.Cards_id != 5 || c.Author_id != 5 {
		t.Error("Fail dans l'update, data non correspondantes", c)
	}
	db.Exec("delete from comments *")
}

func Test_AddComment(t *testing.T) {
	if err := dbPool.InitPool(2, db_open, INFO_CONNECT); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	oldc := &Comment{1, "new super content", 5, 5}
	oldc.Add(dbPool)
	row := db.QueryRow("select id from comments where content = 'new super content'")
	var id int
	row.Scan(&id)
	row = db.QueryRow("select * from comments where id = $1", id)
	c := &Comment{}
	err := row.Scan(&c.Id, &c.Content, &c.Cards_id, &c.Author_id)
	if err != nil || c.Content != "new super content" || c.Cards_id != 5 || c.Author_id != 5 {
		t.Error("Fail dans l'update, data non correspondantes", err)
	}
	db.Exec("delete from comments *")
}

func Test_DelComment(t *testing.T) {
	if err := dbPool.InitPool(2, db_open, INFO_CONNECT); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	var count int
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	db.Exec(`INSERT INTO comments(content, cards_id, author_id)
 VALUES('super content', '10', '10');`)
	row := db.QueryRow("select id from comments where content = 'super content'")
	var id int
	row.Scan(&id)
	c := &Comment{Id: id}
	c.Del(dbPool)
	row = db.QueryRow("select count(*) from comments")
	row.Scan(&count)
	if count != 0 {
		t.Error("delete ne delete pas", count)
		db.Exec("delete from comments *")
	}
}

func Test_GetComment(t *testing.T) {
	if err := dbPool.InitPool(2, db_open, INFO_CONNECT); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	db.Exec(`INSERT INTO comments(content, cards_id, author_id)
 VALUES('super content', '10', '10');`)
	row := db.QueryRow("select id from comments where content = 'super content'")
	var id int
	row.Scan(&id)
	c := &Comment{Id: id}
	err := c.Get(dbPool)
	if err != nil || c.Content != "super content" || c.Cards_id != 10 || c.Author_id != 10 {
		t.Error("le get ne marche pas", c, err)
	}
	db.Exec("delete from comments *")
}
