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

func AddUser(db *sql.DB, u *User) error {
	_, err := db.Exec("INSERT INTO users(name, admin, password, mail, active) VALUES($1, $2, $3, $4, $5);",
		u.Name, u.Admin, u.Password, u.Mail, u.Active)
	return err
}

func DelUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}

func UpdateUser(db *sql.DB, u *User) error {
	_, err := db.Exec("update users set name = $1, admin = $2, password = $3, mail = $4, active = $5 where id = $6",
		u.Name, u.Admin, u.Password, u.Mail, u.Active, u.Id)
	return err
}

// Il faut penser a modifier les fonctions de password pour utiliser des hash plutot que des versions en clair.
func ChangePassword(db *sql.DB, id int, password string) error {
	_, err := db.Exec("update users set password = $1 where id = $2", password, id)
	return err
}

func CheckPassword(db *sql.DB, id int, new_password string) (bool, error) {
	var password string
	var check bool
	row := db.QueryRow("select password from users where id = $1", id)
	err := row.Scan(&password)
	if err != nil && password == new_password {
		check = true
	} else {
		check = false
	}
	return check, err
}

func ChangeUserName(db *sql.DB, id int, name string) error {
	_, err := db.Exec("update users set name = $1 where id = $2", name, id)
	return err
}

func ChangeUserMail(db *sql.DB, id int, mail string) error {
	_, err := db.Exec("update users set mail = $1 where id = $2", mail, id)
	return err
}

func GetNbUsers(db *sql.DB) (int, error) {
	var num int
	row := db.QueryRow("select count(*) from users")
	err := row.Scan(&num)
	return num, err
}

func GetUsersById(db *sql.DB, id int) (*User, error) {
	u := &User{}
	row := db.QueryRow("select * from users where id = $1", id)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return u, err
}

func GetUsersByName(db *sql.DB, name string) (*User, error) {
	u := &User{}
	row := db.QueryRow("select * from users where name = $1", name)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return u, err
}

func changeStateUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set active = $1 where id = $2", state, id)
	return err
}

func ActivateUser(db *sql.DB, id int) error {
	return changeStateUser(db, id, true)
}

func UnactivateUser(db *sql.DB, id int) error {
	return changeStateUser(db, id, false)
}

func changeAdminUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set admin = $1 where id = $2", state, id)
	return err
}

func AdminUser(db *sql.DB, id int) error {
	return changeAdminUser(db, id, true)
}

func UnadminUser(db *sql.DB, id int) error {
	return changeAdminUser(db, id, false)
}
