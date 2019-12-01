package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Methods, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// getID retreive the ID of API request  (http://...../{id})
// r the HTTP Request
//
// return:
// int the id retrived from the Request
// error if an error occurs during id retreival
func getID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	return i, nil
}

//respondJSON create the HTTP answer with JSON content
// w the HTTP writer used to send the response
// status the HTTP status of answer
// payload the message to send
func respondJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

//sendHTTPError send an HTTP answer with an error message
func sendHTTPError(w http.ResponseWriter, errmsg string) {
	enableCors(&w)
	http.Error(w, errmsg, http.StatusInternalServerError)
}

//func getBodyParams(r *http.Request, myg *mg.Group) error {
func getBodyParams(r *http.Request, myg interface{}) error {
	//var myg mg.Group
	if r.Body == nil { //http.Error(w, "Please send a request body", 400)
		log.Println("error get gody")
		return errors.New("Please send a request body")
	}
	err2 := json.NewDecoder(r.Body).Decode(myg)
	if err2 != nil {
		log.Println(r.Body)
		log.Println("error decode")
		//http.Error(w, err2.Error(), 400)
		return err2
	}
	log.Println(myg)
	return nil
}
