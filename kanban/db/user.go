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
	_, err := db.Exec(`INSERT INTO users(name, admin, password, mail, active)
 VALUES($1, $2, $3, $4, $5);`, u.name, u.admin, u.password, u.mail, u.active)
	return err
}

func DelUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}

func GetNbUsers(db *sql.DB) int {
	var num int
	row := db.QueryRow("select count(*) from users")
	if err := row.Scan(&num); err != nil {
		return -1
	}
	return num
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
