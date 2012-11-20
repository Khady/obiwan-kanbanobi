package main

import (
	"testing"
)

func Test_GetNbUsers(t *testing.T) {
	readConf(CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	if _, err := GetNbUsers(dbPool); err != nil {
		t.Error("Impossible de connaitre le nombre d'users")
	}
}

func Test_GetUsersByName(t *testing.T) {
	readConf(CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	name := "super test"
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	u := &User{1, "super test", false, "pass", "user@world.com", true}
	if err := u.GetByName(dbPool); err != nil {
		t.Error("User non existant.", err)
	} else if u.Name != name {
		t.Error("Mauvais name")
	}
	db.Exec("delete from users where name = $1", "super test")
}

func Test_GetUserById(t *testing.T) {
	readConf(CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	u := &User{1, "super test", false, "pass", "user@world.com", true}
	u.GetByName(dbPool)
	id := u.Id
	if err := u.GetById(dbPool); err != nil {
		t.Error("User non existant.", err)
	} else if u.Id != id {
		t.Error("Mauvais id")
	}
	db.Exec("delete from users where name = $1", "super test")
}

func Test_ChangeStateUser(t *testing.T) {
	readConf(CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	u := &User{}
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	row := db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	old_state := u.Active
	if err := changeStateUser(dbPool, u.Id, false); err != nil {
		t.Error("Impossible to change state.", err)
	}
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Active {
		t.Error("The fonc failed to change the state")
	}
	u.Activate(dbPool)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state != u.Active {
		t.Error("The fonc ActivateUser failed to change the state")
	}
	u.Unactivate(dbPool)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Active {
		t.Error("The fonc UnactivateUser failed to change the state")
	}
	db.Exec("delete from users where name = $1", "super test")
}

func Test_ChangeAdminUser(t *testing.T) {
	readConf(CONF_FILE)
	if err := dbPool.InitPool(2, db_open, info_connect_bdd); err != nil {
		t.Error("fail dans l'initpool", err)
	}
	db := dbPool.GetConnection()
	defer dbPool.ReleaseConnection(db)
	u := &User{}
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	row := db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	old_state := u.Admin
	if err := changeAdminUser(dbPool, u.Id, true); err != nil {
		t.Error("Impossible to change state.", err)
	}
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Admin {
		t.Error("The fonc ChangeAdminUser failed to change the state", u.Admin)
	}
	u.Unadmin(dbPool)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state != u.Admin {
		t.Error("The fonc AdminUser failed to change the state")
	}
	u.PutAdmin(dbPool)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Admin {
		t.Error("The fonc UnadminUser failed to change the state")
	}
	db.Exec("delete from users where name = $1", "super test")
}
