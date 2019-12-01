package handlers

import (
	"database/sql"
	"log"
	"net/http"

	mg "phonebook_api/models"
)

//GroupHandler nsdnkjfgnjkd
type GroupHandler struct {
	DBConn *sql.DB `json:"dbConn"`
}

// GetAllGroups return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all groups")
	k, err := mg.GetAllGroups(g.DBConn)
	if err != nil {
		sendHTTPError(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, k)
	}
}

// GetGroupByID return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g GroupHandler) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one group by its ID")
	//var k []*mg.Group
	i, err := getID(r)
	if err != nil {
		sendHTTPError(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	k, err2 := mg.GetGroupByID(g.DBConn, i)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		//http.Error(w, err2.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, k)
	}
}

// CreateGroup return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	log.Println("Create one group")
	var myg mg.Group
	err := getBodyParams(r, &myg)
	log.Println(myg)
	if err != nil {
		//http.Error(w, err.Error(), 400)
		sendHTTPError(w, err.Error())
		return
	}
	g2, err2 := mg.CreateGroup(g.DBConn, &myg)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		//http.Error(w, err2.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, g2)
	}
}

// UpdateGroup return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	log.Println("Update one group")
	i, err := getID(r)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		sendHTTPError(w, err.Error())
		return
	}
	var myg mg.Group
	err2 := getBodyParams(r, &myg)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		http.Error(w, err2.Error(), 400)
		return
	}
	g2, err3 := mg.UpdateGroup(g.DBConn, i, &myg)
	if err3 != nil {
		sendHTTPError(w, err3.Error())
		http.Error(w, err3.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, g2)
	}
}

// DeleteGroup return all the contact groups
// w the HTTP writer used to send the response
// r the HTTP request
func (g GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete one group")
	i, err := getID(r)
	if err != nil {
		sendHTTPError(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err2 := mg.DeleteGroup(g.DBConn, i)
	if err2 != nil {
		sendHTTPError(w, err2.Error())
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	} else {
		respondJSON(w, r, 200, count)
	}
}
