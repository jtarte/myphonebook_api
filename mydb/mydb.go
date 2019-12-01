package mydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//DBInit initialises DB connection
//
//return:
//*sql.DB the DB connection
//error error raised durint the initialization process
func DBInit() (*sql.DB, error) {
	user := os.Getenv("DB_USER")    //phonebook"
	password := os.Getenv("DB_PWD") //phonebook
	dbname := os.Getenv("DB_NAME")  //phonebook
	host := os.Getenv("DB_HOST")    //"localhost"
	port := os.Getenv("DB_PORT")    //5432

	var db *sql.DB
	var err error //error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		log.Println("Error during connection with database")
		log.Println("Review the env variable to fix the error")
		os.Exit(3)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Println("Unable to ping the database")
		os.Exit(4)
	}
	log.Println("Successfully connected!")
	return db, nil
}
