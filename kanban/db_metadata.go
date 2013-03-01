package main

func (c *Metatdata) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec(`UPDATE metadata SET data_value = $1, WHERE id = $4`,
	c.Data_value)
	return err
}

func (c *Metatdata) Add(p *ConnectionPoolWrapper) error {
    db := p.GetConnection()
    defer p.ReleaseConnection(db)
    _, err := db.Exec("INSERT INTO metadata(object_type, object_id, data_key, data_value) VALUES($1, $2, $3, $4);",
	c.Object_type, c.Object_id, c.Data_key, Data_value)
    return err
}

func (c *Metatdata) Del(p *ConnectionPoolWrapper) error {
    db := p.GetConnection()
    defer p.ReleaseConnection(db)
    _, err := db.Exec("DELETE from metadata where id = $1", c.Id)
    return err
}

func (c *Metatdata) Get(p *ConnectionPoolWrapper) error {
    db := p.GetConnection()
    defer p.ReleaseConnection(db)
    row := db.QueryRow("SELECT * FROM metadata WHERE id = $1 AND data_key = $2", c.Id, c.Data_key)
    err := row.Scan(&c.Id, &c.Object_type, &c.Object_id, &c.Data_key, &c.Data_value)
    return err
}
