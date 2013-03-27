package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"errors"
)

func buildNotify(projectId uint32, msg *message.Msg) {
	listOfUsers, err := getUInt32SliceCell(dbPool, "projects", "read", projectId)
	listOfAdmins, errAdmin := getUInt32SliceCell(dbPool, "projects", "admins_id", projectId)
	listOfUsers = append(listOfUsers, listOfAdmins...)
	if err == nil && errAdmin == nil {
		for _, value := range listOfUsers {
			conn, ok := CONNECTION_LIST.ids[value]
			if ok {
				sendKanbanMsg(conn.c, msg)
			}
		}
	}
}


func notifyUsers(msg *message.Msg) error {
	var err error = nil
	switch *msg.Target {
	case message.TARGET_COLUMNS:
		buildNotify(*msg.Columns.ProjectId, msg)
	case message.TARGET_CARDS:
		buildNotify(*msg.Cards.ProjectId, msg)
	case message.TARGET_PROJECTS:
		buildNotify(*msg.Projects.Id, msg)
	// il va falloir regarder comment sont geres les commentaires. Si ce sont oui ou non des metadatas.
	// case message.TARGET_COMMENT:
	// 	comment := &Comment{
	// 		id: *msg.Comments.Id,
	// 		cardsId: *msg.Comments.CardsId,
	// 	}
	// 	buildNotify(*msg.Projects.Id, msg)
	default:
		err = errors.New("Invalid message target")
	}
	return err
}
