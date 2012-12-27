package main

import (
	"testing"
)

func Test_UpdateCard(t *testing.T) {
	readConf(TEST_CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	card := &Card{
		0,
		"First card",
		"awesome content",
		2,
		2,
		[]string{"tag1", "tag2", "tag3"},
		[]int{1, 2, 3},
		1,
		[]int{7, 8, 9},
	}
	if err := card.Add(dbPool); err != nil {
		t.Error(err)
	}
	db.Exec("delete from cards *")
	card.GetScript(dbPool)
}
