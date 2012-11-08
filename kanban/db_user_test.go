package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"testing"
)

func Test_GetNbUsers(t *testing.T) {
	db, _ := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
	defer db.Close()
	if _, err := GetNbUsers(db); err != nil {
		t.Error("Impossible de connaitre le nombre d'users")
	}
}

func Test_GetUserById(t *testing.T) {
	db, _ := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
	defer db.Close()
	id := 1
	u := &User{1, "adm", false, "pass", "user@world.com", true}
	if err := u.GetById(db); err != nil {
		t.Error("User non existant.", err)
	} else if u.Id != id {
		t.Error("Mauvais id")
	}
}

func Test_GetUsersByName(t *testing.T) {
	db, _ := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
	defer db.Close()
	name := "adm"
	u := &User{1, "adm", false, "pass", "user@world.com", true}
	if err := u.GetByName(db); err != nil {
		t.Error("User non existant.", err)
	} else if u.Name != name {
		t.Error("Mauvais name")
	}
}

func Test_ChangeStateUser(t *testing.T) {
	db, _ := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
	defer db.Close()
	u := &User{}
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	row := db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	old_state := u.Active
	if err := changeStateUser(db, u.Id, false); err != nil {
		t.Error("Impossible to change state.", err)
	}
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Active {
		t.Error("The fonc failed to change the state")
	}
	u.Activate(db)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state != u.Active {
		t.Error("The fonc ActivateUser failed to change the state")
	}
	u.Unactivate(db)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Active {
		t.Error("The fonc UnactivateUser failed to change the state")
	}
	db.Exec("delete from users where name = $1", "super test")
}

func Test_ChangeAdminUser(t *testing.T) {
	db, _ := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
	defer db.Close()
	u := &User{}
	db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES('super test', 'false', 'pass', 'user@world.com', 'true');`)
	row := db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	old_state := u.Admin
	if err := changeAdminUser(db, u.Id, true); err != nil {
		t.Error("Impossible to change state.", err)
	}
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Admin {
		t.Error("The fonc ChangeAdminUser failed to change the state", u.Admin)
	}
	u.Unadmin(db)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state != u.Admin {
		t.Error("The fonc AdminUser failed to change the state")
	}
	u.PutAdmin(db)
	row = db.QueryRow("select * from users where name = $1", "super test")
	row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	if old_state == u.Admin {
		t.Error("The fonc UnadminUser failed to change the state")
	}
	db.Exec("delete from users where name = $1", "super test")
}
