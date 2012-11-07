package db

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

type user struct {
	id       int
	name     string
	admin    bool
	password string
	mail     string
	active   bool
}

func AddUser(db *sql.DB, u *user) error {
	_, err := db.Exec("INSERT INTO users(name, admin, password, mail, active) VALUES($1, $2, $3, $4, $5);",
		u.name, u.admin, u.password, u.mail, u.active)
	return err
}

func DelUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}

func UpdateUser(db *sql.DB, u *user) error {
	_, err := db.Exec("update users set name = $1, admin = $2, password = $3, mail = $4, active = $5 where id = $6",
		u.name, u.admin, u.password, u.mail, u.active, u.id)
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

func ChangeName(db *sql.DB, id int, name string) error {
	_, err := db.Exec("update users set name = $1 where id = $2", name, id)
	return err
}

func ChangeMail(db *sql.DB, id int, mail string) error {
	_, err := db.Exec("update users set mail = $1 where id = $2", mail, id)
	return err
}

func GetNbUsers(db *sql.DB) (int, error) {
	var num int
	row := db.QueryRow("select count(*) from users")
	err := row.Scan(&num);
	return num, err
}

func GetUsersById(db *sql.DB, id int) (*user, error) {
	u := &user{}
	row := db.QueryRow("select * from users where id = $1", id)
	err := row.Scan(&u.id, &u.name, &u.admin, &u.password, &u.mail, &u.active)
	return u, err
}

func GetUsersByName(db *sql.DB, name string) (*user, error) {
	u := &user{}
	row := db.QueryRow("select * from users where name = $1", name)
	err := row.Scan(&u.id, &u.name, &u.admin, &u.password, &u.mail, &u.active)
	return u, err
}

func ChangeStateUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set active = $1 where id = $2", state, id)
	return err
}

func ActivateUser(db *sql.DB, id int) error {
	return ChangeStateUser(db, id, true)
}

func UnactivateUser(db *sql.DB, id int) error {
	return ChangeStateUser(db, id, false)
}

func ChangeAdminUser(db *sql.DB, id int, state bool) error {
	_, err := db.Exec("update users set admin = $1 where id = $2", state, id)
	return err
}

func AdminUser(db *sql.DB, id int) error {
	return ChangeAdminUser(db, id, true)
}

func UnadminUser(db *sql.DB, id int) error {
	return ChangeAdminUser(db, id, false)
}

// func Connexion() (*user, error) {
// 	db, err := sql.Open("postgres", "user=kanban password=mdp dbname=kanban")
// 	defer db.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := db.Query("select * from users;")
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := &user{}
// 	for res.Next() {
// 		err = res.Scan(&u.id, &u.name, &u.admin, &u.password, &u.mail, &u.active)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return u, nil
// }
