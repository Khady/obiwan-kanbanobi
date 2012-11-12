package main

type Card struct {
	Id         int
	Name       string
	Content    string
	Column_id  int
	Project_id int
	Tags       string
	Users_id   string
	Scripts_id string
	Write      string
}

func (c *Card) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec(`update cards set name = $1, content = $2, column_id = $3, project_id = $4,
tags = $5, users_id = $6, scripts_id = $7, write = $8 where id = $9`,
		c.Name, c.Content, c.Column_id, c.Project_id, c.Tags, c.Users_id, c.Scripts_id, c.Write, c.Id)
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

func (c *Card) UpdateWrite(p *ConnectionPoolWrapper, write string) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	db = db
	c.Id = c.Id
	write = write
	return nil
}

// Ces deux fonctions vont etre utilisees a la fois pour la gestion des
// scripts et pour la gestions des drois d'ecriture puisque le comportement
// est exactement le meme. Il sagit de modifier correctement la chaine contenant
// une liste d'id et de la retourner
func (c *Card) delIdInCell(cell string) (string, error) {
	return "", nil
}

func (c *Card) addIdInCell(cell string) (string, error) {
	return "", nil
}

// Ajoute des droits de modifications sur la carte a une personne.
// Modifie la chaine deja existante pour y ajouer l'utilisateur correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddWrite(p *ConnectionPoolWrapper, id int) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var wr string
	row := db.QueryRow("select write from cards where id = $1", c.Id)
	if err := row.Scan(&wr); err != nil {
		return err
	}
	// Verifier si id dans wr
	// oui -> ne rien changer
	// non -> le rajouter
	if wr == "" {
	} else {
	}
	return nil
}

// Suppression d'un utilisateur de la chaine de write
// Ne renvoie pas d'erreur si l'utilisateur n'etait pas present
func (c *Card) DelWrite(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	var wr string
	row := db.QueryRow("select write from cards where id = $1", c.Id)
	if err := row.Scan(&wr); err != nil {
		return err
	}
	return nil
}

// Ajoute un script sur la carte.
// Modifie la chaine deja existante pour y ajouer le script correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddScript() {

}

// Suppression d'un script de la chaine de write
// Ne renvoie pas d'erreur si l'script n'etait pas present
func (c *Card) DelScript() {

}
