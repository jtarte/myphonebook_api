package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"phonebook_api/handlers"

	db "phonebook_api/mydb"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main
// entry point of the program
func main() {
	//init the DB connection
	dbconn, err := db.DBInit()
	if err != nil {
		panic(err)
	}
	// defer the release of DB connection
	defer dbconn.Close()
	// define the URL handler
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods(http.MethodGet)
	router.HandleFunc("/health", health).Methods(http.MethodGet)
	initHandlers(router, dbconn)
	router.Use(mux.CORSMethodMiddleware(router))
	// luanch the http server
	myport := os.Getenv("PORT")
	if myport == "" {
		myport = ":8080"
	} else {
		myport = ":" + myport
	}
	log.Printf("Server started on %s \n", myport)
	log.Fatal(http.ListenAndServe(myport, router))
}

// initHandlers initialises the URL handlers
// router the router hosting the handlers
// dbconn the DB connection used to interact with database
func initHandlers(router *mux.Router, dbconn *sql.DB) {
	sub := router.PathPrefix("/api/v1").Subrouter()

	gHandler := handlers.GroupHandler{DBConn: dbconn}
	sub.Methods("GET", "OPTIONS").Path("/groups").HandlerFunc(gHandler.GetAllGroups)
	sub.Methods("GET", "OPTIONS").Path("/groups/{id}").HandlerFunc(gHandler.GetGroupByID)
	sub.Methods("POST").Path("/groups").HandlerFunc(gHandler.CreateGroup)
	sub.Methods("PUT").Path("/groups/{id}").HandlerFunc(gHandler.UpdateGroup)
	sub.Methods("DELETE").Path("/groups/{id}").HandlerFunc(gHandler.DeleteGroup)

	cHandler := handlers.ContactHandler{DBConn: dbconn}
	sub.Methods("GET", "OPTIONS").Path("/contacts").HandlerFunc(cHandler.GetAllContacts)
	sub.Methods("GET", "OPTIONS").Path("/contacts/{id}").HandlerFunc(cHandler.GetContactByID)
	sub.Methods("POST").Path("/contacts").HandlerFunc(cHandler.CreateContact)
	sub.Methods("PUT").Path("/contacts/{id}").HandlerFunc(cHandler.UpdateContact)
	sub.Methods("DELETE").Path("/contacts/{id}").HandlerFunc(cHandler.DeleteContact)
}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(200)
	w.Write([]byte("Hello Jerome"))
}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
