package main

import (
	"strings"
)

func (c *Column) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	tags := strings.Join(c.Tags, " ")
	sid := strings.Join(SString_of_SUInt32(c.Scripts_id), " ")
	write := strings.Join(SString_of_SUInt32(c.Write), " ")
	_, err := db.Exec(`INSERT INTO columns(name, project_id, content, tags, scripts_id, write)
VALUES($1, $2, $3, $4, $5, $6);`,
		c.Name, c.Project_id, c.Content, tags, sid, write)
	return err
}

func (c *Column) Del(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from columns where id = $1", c.Id)
	return err
}

func (c *Column) DelCards(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from cards where column_id = $1", c.Id)
	return err
}

func (c *Column) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	tags := strings.Join(c.Tags, " ")
	sid := strings.Join(SString_of_SUInt32(c.Scripts_id), " ")
	write := strings.Join(SString_of_SUInt32(c.Write), " ")
	_, err := db.Exec(`update columns set name = $1, project_id = $2, content = $3, tags = $4, scripts_id = $5, write = $6 ehere id = $7;`,
		c.Name, c.Project_id, c.Content, tags, sid, write, c.Id)
	return err
}

func (c *Column) Get(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var tags, scripts, write string
	row := db.QueryRow("select * from columns where id = $1", c.Id)
	err := row.Scan(&c.Id, &c.Name, &c.Project_id, &c.Content, &tags, &scripts, &write)
	c.Tags = strings.Split(tags, " ")
	c.Scripts_id = SUInt32_of_SString(strings.Split(scripts, " "))
	c.Write = SUInt32_of_SString(strings.Split(write, " "))
	return err
}

func (c *Column) AddScript(p *ConnectionPoolWrapper) error {
	return nil
}

func (c *Column) DelScript(p *ConnectionPoolWrapper) error {
	return nil
}

func (c *Column) AddTag(p *ConnectionPoolWrapper) error {
	return nil
}

func (c *Column) DelTag(p *ConnectionPoolWrapper) error {
	return nil
}

func (c *Column) GetTags(p *ConnectionPoolWrapper) error {
	return nil
}

func (c *Column) Rename(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update columns set name = $1 where id = $2", c.Name, c.Id)
	return err
}
