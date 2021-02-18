package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	authKey   = "Willem"
	dbDetails = "root@(localhost:5006)/gogenda?parseTime=true"
	db        *sqlx.DB
	format    = "02-01-2006"
)

func getAgenda(w http.ResponseWriter, r *http.Request) {
	selectedItems := []itemStruct{}
	err := db.Select(&selectedItems, "SELECT * FROM agenda_items")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(selectedItems)
}

func addAgendaItem(w http.ResponseWriter, r *http.Request) {
	var bodyValues addItemStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	_, err := db.Query("INSERT INTO agenda_items (name, information, due_date) VALUES (?, ?, ?)", bodyValues.Name, bodyValues.Information, bodyValues.Date)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(204)
}

func getAgendaItems(w http.ResponseWriter, r *http.Request) {
	selectedItems := []itemStruct{}
	query := r.URL.Query()

	switch {
	case query.Get("after") != "" && query.Get("before") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date BETWEEN ? AND ?", query.Get("after"), query.Get("before"))
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(selectedItems)
	case query.Get("date") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date = ?", query.Get("date"))
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(selectedItems)

	default:
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items")
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(selectedItems)
	}
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

	router.HandleFunc("/get_agenda_items", getAgendaItems)
	router.HandleFunc("/add_agenda_items", authenticationCheck(addAgendaItem)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type addItemStruct struct {
	Name        string `db:"name" json:"name"`
	Information string `db:"information" json:"info"`
	Date        string `db:"due_date" json:"date"`
}

type itemStruct struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Information string `db:"information"`
	DueDate     string `db:"due_date"`
	Done        bool   `db:"done"`
}
