package main

import (
	"strings"
)

type cellUpdate func(*Card, *ConnectionPoolWrapper, []int) error
type cellGet func(*Card, *ConnectionPoolWrapper) ([]int, error)

func (c *Card) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	tags := strings.Join(c.Tags, " ")
	sid := strings.Join(SString_of_SUInt32(c.Scripts_id), " ")
	write := strings.Join(SString_of_SUInt32(c.Write), " ")
	_, err := db.Exec(`INSERT INTO cards(name, content, column_id, project_id, tags, user_id, scripts_id, write)
VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
		c.Name, c.Content, c.Column_id, c.Project_id, tags, c.User_id, sid, write)
	return err
}

func (c *Card) Del(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from cards where id = $1", c.Id)
	return err
}

func (c *Card) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	tags := strings.Join(c.Tags, " ")
	sid := strings.Join(SString_of_SUInt32(c.Scripts_id), " ")
	write := strings.Join(SString_of_SUInt32(c.Write), " ")
	_, err := db.Exec(`update cards set name = $1, content = $2, column_id = $3, project_id = $4,
tags = $5, user_id = $6, scripts_id = $7, write = $8 where id = $9`,
		c.Name, c.Content, c.Column_id, c.Project_id, tags, c.User_id, sid, write, c.Id)
	return err
}

func (c *Card) Get(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var tags, scripts, write string
	row := db.QueryRow("select * from cards where id = $1", c.Id)
	err := row.Scan(&c.Id, &c.Name, &c.Content, &c.Column_id, &c.Project_id, &tags, &c.User_id,
		&scripts, &write)
	c.Tags = strings.Split(tags, ",")
	c.Scripts_id = SUInt32_of_SString(strings.Split(scripts, ","))
	c.Write = SUInt32_of_SString(strings.Split(write, ","))
	return err
}

func (c *Card) ChangeColumn_id(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update cards set column_id = $1 where id = $2", c.Column_id, c.Id)
	return err
}

func (c *Card) ChangeName(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update cards set name = $1 where id = $2", c.Name, c.Id)
	return err
}

func (c *Card) ChangeContent(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update cards set content = $1 where id = $2", c.Content, c.Id)
	return err
}

func (c *Card) UpdateTags(p *ConnectionPoolWrapper, tags []string) error {
	return updateStringSliceCell(p, "cards", "tags", tags, c.Id)
}

func (c *Card) GetTags(p *ConnectionPoolWrapper) ([]string, error) {
	return getStringSliceCell(p, "cards", "tags", c.Id)
}

func (c *Card) UpdateWrite(p *ConnectionPoolWrapper, write []uint32) error {
	return updateUInt32SliceCell(p, "cards", "write", write, c.Id)
}

func (c *Card) GetWrite(p *ConnectionPoolWrapper) ([]uint32, error) {
	return getUInt32SliceCell(p, "cards", "write", c.Id)
}

func (c *Card) UpdateScript(p *ConnectionPoolWrapper, script []uint32) error {
	return updateUInt32SliceCell(p, "cards", "scripts_id", script, c.Id)
}

func (c *Card) GetScript(p *ConnectionPoolWrapper) ([]uint32, error) {
	return getUInt32SliceCell(p, "cards", "scripts_id", c.Id)
}

// Ajoute des droits de modifications sur la carte a une personne.
// Modifie la chaine deja existante pour y ajouer l'utilisateur correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddWrite(p *ConnectionPoolWrapper, id uint32) error {
	return addIdInCell(p, id, c.Id, "cards", "write")
}

// Suppression d'un utilisateur de la chaine de write
// Ne renvoie pas d'erreur si l'utilisateur n'etait pas present
func (c *Card) DelWrite(p *ConnectionPoolWrapper, id uint32) error {
	return delIdInCell(p, id, c.Id, "cards", "write")
}

// Ajoute un script sur la carte.
// Modifie la chaine deja existante pour y ajouer le script correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddScript(p *ConnectionPoolWrapper, id uint32) error {
	return addIdInCell(p, id, c.Id, "cards", "scripts_id")
}

// Suppression d'un script de la chaine de write
// Ne renvoie pas d'erreur si l'script n'etait pas present
func (c *Card) DelScript(p *ConnectionPoolWrapper, id uint32) error {
	return delIdInCell(p, id, c.Id, "cards", "scripts_id")

}

func (c *Card)GetLastCardWithName(p* ConnectionPoolWrapper) error {
        db := p.GetConnection()
        defer p.ReleaseConnection(db)
	row := db.QueryRow("select max(id) from cards where name= $1", c.Name)
        err := row.Scan(&c.Id)
        if (err != nil) {
		println("lol")
        }
        return err
}
