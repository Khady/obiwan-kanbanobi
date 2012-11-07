package db

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

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

func UpdateCard(db *sql.DB, c *Card) error {
	_, err := db.Exec(`update cards set name = $1, content = $2, column_id = $3, project_id = $4,
tags = $5, users_id = $6, scripts_id = $7, write = $8 where id = $9`,
		c.Name, c.Content, c.Column_id, c.Project_id, c.Tags, c.Users_id, c.Scripts_id, c.Write, c.Id)
	return err
}

func ChangeCardName(db *sql.DB, id int, name string) error {
	_, err := db.Exec("update cards set name = $1 where id = $2", name, id)
	return err
}

func ChangeCardContent(db *sql.DB, id int, content string) error {
	_, err := db.Exec("update cards set content = $1 where id = $2", content, id)
	return err
}

// Ajoute des droits de modifications sur la carte a une personne.
// Modifie la chaine deja existante pour y ajouer l'utilisateur correctement
// Creation de la chaine si elle n'existait pas deja
func AddWriteCard() {

}

// Suppression d'un utilisateur de la chaine de write
// Ne renvoie pas d'erreur si l'utilisateur n'etait pas present
func DelWriteCard() {

}
