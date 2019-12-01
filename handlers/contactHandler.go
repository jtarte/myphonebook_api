package handlers

import (
	"database/sql"
	"log"
	"net/http"

	mg "phonebook_api/models"
)

//ContactHandler nsdnkjfgnjkd
type ContactHandler struct {
	DBConn *sql.DB `json:"dbConn"`
}

// GetAllContacts return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all contacts")
	k, err := mg.GetAllContacts(g.DBConn)
	if err != nil {
		sendHTTPError(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, k)
	}
}

// GetContactByID return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g ContactHandler) GetContactByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one contact by its ID")
	//var k []*mg.Group
	i, err := getID(r)
	if err != nil {
		sendHTTPError(w, err.Error())
		return
	}
	k, err2 := mg.GetContactByID(g.DBConn, i)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
	} else {
		respondJSON(w, r, 200, k)
	}
}

// CreateContact return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	log.Println("Create one contact")
	var myg mg.Contact
	err := getBodyParams(r, &myg)
	if err != nil {
		log.Println("error in get param")
		sendHTTPError(w, err.Error())
		return
	}
	g2, err2 := mg.CreateContact(g.DBConn, &myg)
	if err2 != nil {
		log.Println("error in create contact model")
		sendHTTPError(w, err2.Error())

	} else {
		respondJSON(w, r, 200, g2)
	}
}

// UpdateContact return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	log.Println("Update one contact")
	i, err := getID(r)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		sendHTTPError(w, err.Error())
		return
	}
	var myg mg.Contact
	err2 := getBodyParams(r, &myg)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		//http.Error(w, err2.Error(), 400)
		return
	}
	g2, err3 := mg.UpdateContact(g.DBConn, i, &myg)
	if err3 != nil {
		//http.Error(w, err3.Error(), http.StatusInternalServerError)
		sendHTTPError(w, err3.Error())
	} else {
		respondJSON(w, r, 200, g2)
	}
}

// DeleteContact return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete one contact")
	i, err := getID(r)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		sendHTTPError(w, err.Error())
		return
	}
	count, err2 := mg.DeleteContact(g.DBConn, i)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		//http.Error(w, err2.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, count)
	}
}
