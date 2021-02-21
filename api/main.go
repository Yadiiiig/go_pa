package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	authKey         = "Willem"
	dbDetails       = "root@(localhost:5006)/gogenda?parseTime=true" //for localhost connection
	dbDetailsDocker = "root@(db:3306)/gogenda?parseTime=true"        // db is name of docker container, 3006 is default port
	db              *sqlx.DB
	format          = "02-01-2006"
	blocked         []string
	useDocker       = false
)

func main() {
	var err error
	router := mux.NewRouter().StrictSlash(true)

	if useDocker {
		db, err = sqlx.Connect("mysql", dbDetailsDocker)
	} else {
		db, err = sqlx.Connect("mysql", dbDetails)
	}
	if err != nil {
		panic(err)
	}

	initBlockedIPs()

	// Agenda routes
	router.HandleFunc("/get_agenda_items", authenticationCheck(getAgendaItems)).Methods("GET")
	router.HandleFunc("/add_agenda_items", authenticationCheck(addAgendaItem)).Methods("POST")
	router.HandleFunc("/delete_agenda_item", authenticationCheck(deleteAgendaItem)).Methods("DELETE")

	// Notes routes
	router.HandleFunc("/get_notes", authenticationCheck(getNotes)).Methods("GET")
	router.HandleFunc("/add_note", authenticationCheck(addNote)).Methods("POST")
	router.HandleFunc("/update_note", authenticationCheck(updateNote)).Methods("PATCH")
	router.HandleFunc("/delete_note", authenticationCheck(deleteNote)).Methods("DELETE")

	// Classes routes
	router.HandleFunc("/get_classes", authenticationCheck(getClasses)).Methods("GET")
	router.HandleFunc("/update_class", authenticationCheck(updateClass)).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
