package main

import (
	"encoding/json"
	"net/http"
)

func addAgendaItem(w http.ResponseWriter, r *http.Request) {
	var bodyValues addItemStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	_, err := db.Query("INSERT INTO agenda_items (name, information, due_date) VALUES (?, ?, ?)", bodyValues.Name, bodyValues.Information, bodyValues.Date)
	if databaseError(w, err) {
		return
	}

	w.WriteHeader(204)
}

func getAgendaItems(w http.ResponseWriter, r *http.Request) {
	selectedItems := []itemStruct{}
	query := r.URL.Query()

	switch {
	case query.Get("after") != "" && query.Get("before") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date BETWEEN ? AND ?", query.Get("after"), query.Get("before"))
		if databaseError(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)

	case query.Get("date") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date = ?", query.Get("date"))
		if databaseError(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)

	case query.Get("id") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE id = ?", query.Get("id"))
		if databaseError(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)

	default:
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items")
		if databaseError(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)

	}
}

func deleteAgendaItem(w http.ResponseWriter, r *http.Request) {
	var bodyValues deleteItemStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	_, err := db.Query("DELETE FROM agenda_items WHERE id = ?", bodyValues.ID)
	if databaseError(w, err) {
		return
	}

	json.NewEncoder(w).Encode(bodyValues.ID)
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

type deleteItemStruct struct {
	ID int `db:"id" json:"id"`
}
