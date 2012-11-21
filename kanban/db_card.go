package main

import (
	"sort"
	"strings"
)

type cellUpdate func(*Card, *ConnectionPoolWrapper, []int) error
type cellGet func(*Card, *ConnectionPoolWrapper) ([]int, error)

func (c *Card) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	tags := strings.Join(c.Tags, " ")
	uid := strings.Join(SString_of_SInt(c.Users_id), " ")
	sid := strings.Join(SString_of_SInt(c.Scripts_id), " ")
	write := strings.Join(SString_of_SInt(c.Write), " ")
	_, err := db.Exec(`INSERT INTO cards(name, content, column_id, project_id, tags, users_id, scripts_id, write)
VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
		c.Name, c.Content, c.Column_id, c.Project_id, tags, uid, sid, write)
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
	uid := strings.Join(SString_of_SInt(c.Users_id), " ")
	sid := strings.Join(SString_of_SInt(c.Scripts_id), " ")
	write := strings.Join(SString_of_SInt(c.Write), " ")
	_, err := db.Exec(`update cards set name = $1, content = $2, column_id = $3, project_id = $4,
tags = $5, users_id = $6, scripts_id = $7, write = $8 where id = $9`,
		c.Name, c.Content, c.Column_id, c.Project_id, tags, uid, sid, write, c.Id)
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

// PRESQUE TOUTES LES FONCTIONS QUI SUIVENT PEUVENT ETRE UTILISEES DE MANIERE TRES PROCHE
// PAR LES COLONNES ET LES PROJETS, IL FAUDRA PENSER A BIEN FACTORISER

func (c *Card) UpdateTags(p *ConnectionPoolWrapper, tags []string) error {
	return updateStringSliceCell(p, "cards", "tags", tags, c.Id)
}

func (c *Card) GetTags(p *ConnectionPoolWrapper) ([]string, error) {
	return getStringSliceCell(p, "cards", "tags", c.Id)
}

func (c *Card) UpdateWrite(p *ConnectionPoolWrapper, write []int) error {
	return updateIntSliceCell(p, "cards", "write", write, c.Id)
}

func (c *Card) GetWrite(p *ConnectionPoolWrapper) ([]int, error) {
	return getIntSliceCell(p, "cards", "write", c.Id)
}

func (c *Card) UpdateScript(p *ConnectionPoolWrapper, script []int) error {
	return updateIntSliceCell(p, "cards", "scripts_id", script, c.Id)
}

func (c *Card) GetScript(p *ConnectionPoolWrapper) ([]int, error) {
	return getIntSliceCell(p, "cards", "scripts_id", c.Id)
}

// Ces deux fonctions vont etre utilisees a la fois pour la gestion des
// scripts et pour la gestions des drois d'ecriture puisque le comportement
// est exactement le meme. Il sagit de modifier correctement la chaine contenant
// une liste d'id et de la retourner

func (c *Card) delIdInCell(p *ConnectionPoolWrapper, id int, cg cellGet, cu cellUpdate) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	cell_ids, err := cg(c, dbPool)
	if err != nil {
		return err
	}
	id_idx := sort.Search(len(cell_ids), func(i int) bool { return cell_ids[i] == id }) // C'est pas un peu crade ce search niveau complexite ?
	if id_idx != len(cell_ids) {
		new_ids := []int{}
		new_ids = append(new_ids, cell_ids[:id_idx]...)
		new_ids = append(new_ids, cell_ids[id_idx+1:]...)
		return cu(c, p, new_ids)
	}
	return err
}

func (c *Card) addIdInCell(p *ConnectionPoolWrapper, id int, cg cellGet, cu cellUpdate) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	cell_ids, err := cg(c, p)
	if err != nil {
		return err
	}
	id_idx := sort.Search(len(cell_ids), func(i int) bool { return cell_ids[i] == id }) // C'est pas un peu crade ce search niveau complexite ?
	if id_idx == len(cell_ids) {
		cell_ids = append(cell_ids, id)
		err = cu(c, p, cell_ids)
	}
	return err
}

// Ajoute des droits de modifications sur la carte a une personne.
// Modifie la chaine deja existante pour y ajouer l'utilisateur correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddWrite(p *ConnectionPoolWrapper, id int) error {
	return c.addIdInCell(p, id, (*Card).GetWrite, (*Card).UpdateWrite)
}

// Suppression d'un utilisateur de la chaine de write
// Ne renvoie pas d'erreur si l'utilisateur n'etait pas present
func (c *Card) DelWrite(p *ConnectionPoolWrapper, id int) error {
	return c.delIdInCell(p, id, (*Card).GetWrite, (*Card).UpdateWrite)
}

// Ajoute un script sur la carte.
// Modifie la chaine deja existante pour y ajouer le script correctement
// Creation de la chaine si elle n'existait pas deja
func (c *Card) AddScript(p *ConnectionPoolWrapper, id int) error {
	return c.addIdInCell(p, id, (*Card).GetScript, (*Card).UpdateScript)
}

// Suppression d'un script de la chaine de write
// Ne renvoie pas d'erreur si l'script n'etait pas present
func (c *Card) DelScript(p *ConnectionPoolWrapper, id int) error {
	return c.delIdInCell(p, id, (*Card).GetScript, (*Card).UpdateScript)

}
