package models

import (
	"database/sql"
)

// Group definition of contact group
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetAllGroups return a tab with all contact groups
// db the DB connexion
//
// return:
// []*Group an array with all the groups records
// error the error that has been raised during the select action (nil if all is OK)
func GetAllGroups(db *sql.DB) ([]*Group, error) {
	rows, err := db.Query("SELECT id, name FROM mygroup ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	k := make([]*Group, 0)
	for rows.Next() {
		mg := new(Group)
		err = rows.Scan(&mg.ID, &mg.Name)
		if err != nil {
			return nil, err
		}
		k = append(k, mg)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return k, nil
}

//GetGroupByID gets a Group by its ID
// db the DB connexion
// sid id of the target group
//
// return:
// *Group the selected group object
// error the error that has been raised during the select action (nil if all is OK)
func GetGroupByID(db *sql.DB, sid int) (*Group, error) {
	row := db.QueryRow("SELECT id, name FROM mygroup where id=$1", sid)
	k := new(Group)
	err := row.Scan(&k.ID, &k.Name)
	if err != nil {
		return nil, err
	}
	return k, nil
}

//CreateGroup creates a new Group record
// db the DB connexion
//
// return:
// *Group the created group object
// error the error that has been raised during the insert action (nil if all is OK)
func CreateGroup(db *sql.DB, g *Group) (*Group, error) {
	id := 0
	name := (*g).Name
	err := db.QueryRow("INSERT INTO mygroup (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return nil, err
	}
	g.ID = id
	return g, nil
}

//UpdateGroup updates a record on database
// db the DB connexion
// sid the id of the record to be deleted
// g the values (Group object) that should be used to update the records
//
// return:
// *Group the updated group object
// error the error that has been raised during the update action (nil if all is OK)
func UpdateGroup(db *sql.DB, sid int, g *Group) (*Group, error) {
	id := 0
	name := (*g).Name
	err := db.QueryRow("UPDATE mygroup SET id = $1 , name = $2 where id = $1 RETURNING id, name", sid, name).Scan(&id, &name)
	if err != nil {
		return nil, err
	}
	g.ID = id
	g.Name = name
	return g, nil
}

//DeleteGroup deletes a group from database
// db the DB connexion
//sid the id of the record to be deleted
//
//return:
//int64 the number of record that has been deleted (0 or 1 if all is OK, -1 if an error is thrown)
//error the error that has been raised during the delete action (nil if all is OK)
func DeleteGroup(db *sql.DB, sid int) (int64, error) {
	var query = "DELETE FROM mygroup where id=$1"
	return delete(db, query, sid)
}
