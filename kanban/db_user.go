package main

type User struct {
	Id       int
	Name     string
	Admin    bool
	Password string
	Mail     string
	Active   bool
}

func changeStateUser(p *ConnectionPoolWrapper, id int, state bool) error {
	db := p.GetConnection()
	_, err := db.Exec("update users set active = $1 where id = $2", state, id)
	p.ReleaseConnection(db)
	return err
}

func changeAdminUser(p *ConnectionPoolWrapper, id int, state bool) error {
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
	_, err := db.Exec("update users set name = $1, admin = $2, password = $3, mail = $4, active = $5 where id = $6",
		u.Name, u.Admin, u.Password, u.Mail, u.Active, u.Id)
	return err
}

// Il faut penser a modifier les fonctions de password pour utiliser des hash plutot que des versions en clair.
func (u *User) ChangePassword(p *ConnectionPoolWrapper, password string) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update users set password = $1 where id = $2", password, u.Id)
	return err
}

func (u *User) CheckPassword(p *ConnectionPoolWrapper, new_password string) (bool, error) {
	var password string
	var check bool
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select password from users where id = $1", u.Id)
	err := row.Scan(&password)
	if err != nil && password == new_password {
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
