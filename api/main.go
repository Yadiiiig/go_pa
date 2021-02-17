package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	authKey   = "Willem"
	dbDetails = "root@(localhost:5006)/gogenda?parseTime=true"
	db        *sqlx.DB
)

func addAgendaItem(w http.ResponseWriter, r *http.Request) {
	var bodyValues addItemStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	parsedTime, err := time.Parse("02-01-2006", bodyValues.Date)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	_, err = db.Query("INSERT INTO agenda_items (name, information, due_date) VALUES (?, ?, ?)", bodyValues.Name, bodyValues.Information, parsedTime)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(204)
}

func authenticationCheck(request http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "" {
			if auth == authKey {
				request(w, r)
			} else {
				w.WriteHeader(403)
				json.NewEncoder(w).Encode("What are you trying to accomplish?")
			}
		} else {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode("What are you trying to accomplish?")
		}
	})
}

func main() {
	var err error
	router := mux.NewRouter().StrictSlash(true)

	db, err = sqlx.Connect("mysql", dbDetails)
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/add_agenda_items", authenticationCheck(addAgendaItem)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type addItemStruct struct {
	Name        string `db:"name" json:"name"`
	Information string `db:"information" json:"info"`
	Date        string `db:"due_date" json:"date"`
}
