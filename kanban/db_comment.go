package main

func (c *Comment) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec(`update comments set content = $1, cards_id = $2, author_id = $3 where id = $4`,
		c.Content, c.Cards_id, c.Author_id, c.Id)
	return err
}

func (c *Comment) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("INSERT INTO comments(content, cards_id, author_id, comment_date) VALUES($1, $2, $3, $4);",
		c.Content, c.Cards_id, c.Author_id, c.Comment_date)
	return err
}

func (c *Comment) Del(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from comments where id = $1", c.Id)
	return err
}

func (c *Comment) Get(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from comments where id = $1", c.Id)
	err := row.Scan(&c.Id, &c.Content, &c.Cards_id, &c.Author_id, &c.Comment_date)
	return err
}

func (c *Comment) CountForCard(p *ConnectionPoolWrapper) (int, error) {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var count int
	row := db.QueryRow("select count(*) from comments where cards_id = $1", c.Cards_id)
	err := row.Scan(&count)
	return count, err
}
