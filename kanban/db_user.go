package main

import "strconv"

func changeStateUser(p *ConnectionPoolWrapper, id uint32, state bool) error {
	db := p.GetConnection()
	_, err := db.Exec("update users set active = $1 where id = $2", state, id)
	p.ReleaseConnection(db)
	return err
}

func changeAdminUser(p *ConnectionPoolWrapper, id uint32, state bool) error {
	db := p.GetConnection()
	_, err := db.Exec("update users set admin = $1 where id = $2", state, id)
	p.ReleaseConnection(db)
	return err
}

func GetNbUsers(p *ConnectionPoolWrapper) (int, error) {
	var num int
	db := p.GetConnection()
	row := db.QueryRow("select count(*) from users")
	err := row.Scan(&num)
	p.ReleaseConnection(db)
	return num, err
}

func (u *User) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("INSERT INTO users(name, admin, password, mail, active) VALUES($1, $2, $3, $4, $5);",
		u.Name, u.Admin, u.Password, u.Mail, u.Active)
	return err
}

func (u *User) Del(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from users where id = $1", u.Id)
	return err
}

func (u *User) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update users set name = $1, mail = $2, active = $3 where id = $4",
		u.Name, u.Mail, u.Active, u.Id)
	return err
}

// Il faut penser a modifier les fonctions de password pour utiliser des hash plutot que des versions en clair.
func (u *User) ChangePassword(p *ConnectionPoolWrapper, password string) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update users set password = $1 where id = $2", password, u.Id)
	return err
}

// Pourquoi passer un password en argument alors qu'il serait possible d'utiliser le password de l'User sur lequel -> surement pour pouvoir faire la verif que le password a bien changer par rapport a l'ancien (row scan -> if).
// la methode est appliquee ?
func (u *User) CheckPassword(p *ConnectionPoolWrapper, new_password string) (bool, error) {
	var password string
	var check bool
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select password from users where id = $1", u.Id)
	err := row.Scan(&password)
	if err == nil && password == new_password {
		check = true
	} else {
		check = false
	}
	return check, err
}

func (u *User) ChangeName(p *ConnectionPoolWrapper, name string) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update users set name = $1 where id = $2", name, u.Id)
	return err
}

func (u *User) ChangeMail(p *ConnectionPoolWrapper, mail string) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update users set mail = $1 where id = $2", mail, u.Id)
	return err
}

func (u *User) GetById(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from users where id = $1", u.Id)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return err
}

func (u *User) GetByName(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from users where name = $1", u.Name)
	err := row.Scan(&u.Id, &u.Name, &u.Admin, &u.Password, &u.Mail, &u.Active)
	return err
}

func (u *User) Activate(p *ConnectionPoolWrapper) error {
	return changeStateUser(p, u.Id, true)
}

func (u *User) Unactivate(p *ConnectionPoolWrapper) error {
	return changeStateUser(p, u.Id, false)
}

func (u *User) PutAdmin(p *ConnectionPoolWrapper) error {
	return changeAdminUser(p, u.Id, true)
}

func (u *User) Unadmin(p *ConnectionPoolWrapper) error {
	return changeAdminUser(p, u.Id, false)
}

func (u *User) GetProjectByUserId(p *ConnectionPoolWrapper) ([]Project, error) {
	var tab []Project
	var t Project

	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row, err := db.Query("SELECT id, name, content FROM projects WHERE read LIKE '%," + strconv.FormatUint(uint64(u.Id), 10) + ",%'")
	if err != nil {
		return tab, err
	}
	for row.Next() {
		err = row.Scan(&t.Id, &t.Name, &t.Content)
		if err != nil {
			return tab, err
		}
		tab = append(tab, t)
	}

	return tab, nil
}

func (u *User) GetAdminById(p *ConnectionPoolWrapper, id uint32) (bool, error) {
	var admin bool

	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select admin from users where id = $1", id)
	err := row.Scan(&admin)
	return admin, err
}
