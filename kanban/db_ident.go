package main

func (s *Session) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	id := &Session{}
	row := db.QueryRow("select * from session where user_id = $1", s.User_id)
	err := row.Scan(&id.Id, &id.User_id, &id.Ident_date, &id.Session_key)
	if err != nil { // there is an error if there is no row with the user login in the DB. So add it.
		_, err = db.Exec("INSERT INTO sessions(user_id, ident_date, session_key) VALUES($1, $2, $3);",
			s.User_id, s.Ident_date, s.Session_key)
	}
	return err
}

func (s *Session) GetUserSessionByName(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	id := &Session{}
	row := db.QueryRow("select * from session where user_id = $1", s.User_id)
	err := row.Scan(&id.Id, &id.User_id, &id.Ident_date, &id.Session_key)
	return err
}

func (s *Session) GetUserSessionById(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	id := &Session{}
	row := db.QueryRow("select * from session where id = $1", s.Id)
	err := row.Scan(&id.Id, &id.User_id, &id.Ident_date, &id.Session_key)
	return err
}

func (s *Session) GetUserSessionBySessionId(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	id := &Session{}
	row := db.QueryRow("select * from session where session_key = $1", s.Session_key)
	err := row.Scan(&id.Id, &id.User_id, &id.Ident_date, &id.Session_key)
	return err
}

func (s *Session) DelByUserName(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	id := &Session{}
	row := db.QueryRow("select * from session where user_id = $1", s.User_id)
	err := row.Scan(&id.Id, &id.User_id, &id.Ident_date, &id.Session_key)
	if err == nil {
		_, err = db.Exec("delete from session where id = $1", id.Id)
	}
	return err
}
