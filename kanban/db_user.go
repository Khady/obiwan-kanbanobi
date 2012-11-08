package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

type User struct {
	Id       int
	Name     string
	Admin    bool
	Password string
	Mail     string
	Active   bool
}

func changeStateUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set active = $1 where id = $2", state, id)
	return err
}

func changeAdminUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set admin = $1 where id = $2", state, id)
	return err
}

func GetNbUsers(db *sql.DB) (int, error) {
	var num int
	row := db.QueryRow("select count(*) from users")
	err := row.Scan(&num)
	return num, err
}

func (u *User) Add(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users(name, admin, password, mail, active) VALUES($1, $2, $3, $4, $5);",
		u.Name, u.Admin, u.Password, u.Mail, u.Active)
	return err
}

func (u *User) Del(db *sql.DB) error {
	_, err := db.Exec("delete from users where id = $1", u.Id)
	return err
}

func (u *User) Update(db *sql.DB) error {
	_, err := db.Exec("update users set name = $1, admin = $2, password = $3, mail = $4, active = $5 where id = $6",
		u.Name, u.Admin, u.Password, u.Mail, u.Active, u.Id)
	return err
}

// Il faut penser a modifier les fonctions de password pour utiliser des hash plutot que des versions en clair.
func (u *User) ChangePassword(db *sql.DB, password string) error {
	_, err := db.Exec("update users set password = $1 where id = $2", password, u.Id)
	return err
}

func (u *User) CheckPassword(db *sql.DB, new_password string) (bool, error) {
	var password string
	var check bool
	row := db.QueryRow("select password from users where id = $1", u.Id)
	err := row.Scan(&password)
	if err != nil && password == new_password {
		check = true
	} else {
		check = false
	}
	return check, err
}

func (u *User) ChangeName(db *sql.DB, name string) error {
	_, err := db.Exec("update users set name = $1 where id = $2", name, u.Id)
	return err
}

func (u *User) ChangeMail(db *sql.DB, mail string) error {
	_, err := db.Exec("update users set mail = $1 where id = $2", mail, u.Id)
	return err
}

func (u *User) GetById(db *sql.DB) error {
	row := db.QueryRow("select * from users where id = $1", u.Id)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return err
}

func (u *User) GetByName(db *sql.DB) error {
	row := db.QueryRow("select * from users where name = $1", u.Name)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return err
}

func (u *User) Activate(db *sql.DB) error {
	return changeStateUser(db, u.Id, true)
}

func (u *User) Unactivate(db *sql.DB) error {
	return changeStateUser(db, u.Id, false)
}

func (u *User) PutAdmin(db *sql.DB) error {
	return changeAdminUser(db, u.Id, true)
}

func (u *User) Unadmin(db *sql.DB) error {
	return changeAdminUser(db, u.Id, false)
}
