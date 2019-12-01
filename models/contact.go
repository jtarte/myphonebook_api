package models

import (
	"database/sql"
	"log"
)

// Contact definition of contact
type Contact struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	GroupID     int    `json:"groupid"`
	GroupName   string `json:"groupname"`
}

// GetAllContacts return a tab with all contact groups
// db the DB connexion
//
// return:
// []*Contact an array with all the contact records
// error the error that has been raised during the select action (nil if all is OK)
func GetAllContacts(db *sql.DB) ([]*Contact, error) {
	log.Println("test")
	rows, err := db.Query("SELECT p.id,p.firstname,p.name,p.email,p.phonenumber,p.gr_id, g.name FROM person p LEFT OUTER JOIN mygroup g ON (p.gr_id = g.id);")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	k := make([]*Contact, 0)
	for rows.Next() {
		mg := new(Contact)
		var rid, rgid sql.NullInt64
		var rf, rn, re, rp, rgn sql.NullString
		err = rows.Scan(&rid, &rf, &rn, &re, &rp, &rgid, &rgn)
		if err != nil {
			return nil, err
		}
		mg.ID = int((&rid).Int64)
		mg.Firstname = (&rf).String
		mg.Name = (&rn).String
		mg.Email = (&re).String
		mg.PhoneNumber = (&rp).String
		mg.GroupID = int((&rgid).Int64)
		mg.GroupName = (&rgn).String

		k = append(k, mg)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return k, nil
}

//GetContactByID gets a contact by its ID
// db the DB connexion
// sid id of the target group
//
// return:
// *Contact the selected group object
// error the error that has been raised during the select action (nil if all is OK)
func GetContactByID(db *sql.DB, sid int) (*Contact, error) {
	row := db.QueryRow("SELECT p.id,p.firstname,p.name,p.email,p.phonenumber,p.gr_id, g.name FROM person p LEFT OUTER JOIN mygroup g ON (p.gr_id = g.id) where p.id=$1", sid)
	mg := new(Contact)
	var rid, rgid sql.NullInt64
	var rf, rn, re, rp, rgn sql.NullString
	err := row.Scan(&rid, &rf, &rn, &re, &rp, &rgid, &rgn)
	if err != nil {
		return nil, err
	}
	mg.ID = int((&rid).Int64)
	mg.Firstname = (&rf).String
	mg.Name = (&rn).String
	mg.Email = (&re).String
	mg.PhoneNumber = (&rp).String
	mg.GroupID = int((&rgid).Int64)
	mg.GroupName = (&rgn).String
	return mg, nil
}

//CreateContact creates a new Group record
// db the DB connexion
//
// return:
// *Contact the created group object
// error the error that has been raised during the insert action (nil if all is OK)
func CreateContact(db *sql.DB, g *Contact) (*Contact, error) {
	var id sql.NullInt64
	log.Println(g)
	//name := (*g).Name
	var err error
	if (*g).GroupID == 0 {
		log.Println("gr id =0")
		err = db.QueryRow("INSERT INTO person (firstname, name, email, phonenumber) VALUES ($1,$2,$3,$4) RETURNING id", (*g).Firstname, (*g).Name, (*g).Email, (*g).PhoneNumber).Scan(&id)

	} else {
		log.Println("gr id !=0")
		err = db.QueryRow("INSERT INTO person (firstname, name, email, phonenumber, gr_id) VALUES ($1,$2,$3,$4,$5) RETURNING id", (*g).Firstname, (*g).Name, (*g).Email, (*g).PhoneNumber, (*g).GroupID).Scan(&id)

	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	g.ID = int((&id).Int64)
	return g, nil
}

//UpdateContact updates a record on database
// db the DB connexion
// sid the id of the record to be deleted
// g the values (Group object) that should be used to update the records
//
// return:
// *Contact the updated group object
// error the error that has been raised during the update action (nil if all is OK)
func UpdateContact(db *sql.DB, sid int, g *Contact) (*Contact, error) {
	var id sql.NullInt64
	var fn sql.NullString
	var name sql.NullString
	var em sql.NullString
	var ph sql.NullString
	var grid sql.NullInt64

	var err error
	if (*g).GroupID == 0 {
		log.Println("gr id =0")
		err = db.QueryRow("UPDATE person SET id = $1 , firstname = $2, name = $3, email = $4, phonenumber = $5 where id = $1 RETURNING id, firstname, name, email, phonenumber, gr_id", sid, (*g).Firstname, (*g).Name, (*g).Email, (*g).PhoneNumber).Scan(&id, &fn, &name, &em, &ph, &grid)
	} else {
		log.Println("gr id !=0")
		err = db.QueryRow("UPDATE person SET id = $1 , firstname = $2, name = $3, email = $4, phonenumber = $5, gr_id = $6 where id = $1 RETURNING id, firstname, name, email, phonenumber, gr_id", sid, (*g).Firstname, (*g).Name, (*g).Email, (*g).PhoneNumber, (*g).GroupID).Scan(&id, &fn, &name, &em, &ph, &grid)
	}
	if err != nil {
		return nil, err
	}
	g.ID = int((&id).Int64)
	g.Firstname = (&fn).String
	g.Name = (&name).String
	g.Email = (&em).String
	g.PhoneNumber = (&ph).String
	g.GroupID = int((&grid).Int64)
	return g, nil
}

//DeleteContact deletes a group from database
// db the DB connexion
//sid the id of the record to be deleted
//
//return:
//int64 the number of record that has been deleted (0 or 1 if all is OK, -1 if an error is thrown)
//error the error that has been raised during the delete action (nil if all is OK)
func DeleteContact(db *sql.DB, sid int) (int64, error) {
	var query = "DELETE FROM person where id=$1"
	return delete(db, query, sid)
}
