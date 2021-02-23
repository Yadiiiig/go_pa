package main

import (
	"encoding/json"
	"net/http"
)

func getAgendaItems(w http.ResponseWriter, r *http.Request) {
	selectedItems := []itemStruct{}
	query := r.URL.Query()

	switch {
	case query.Get("after") != "" && query.Get("before") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date BETWEEN ? AND ?", query.Get("after"), query.Get("before"))
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)

	case query.Get("date") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE due_date = ?", query.Get("date"))
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)

	case query.Get("id") != "":
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)

	default:
		err := db.Select(&selectedItems, "SELECT * FROM agenda_items")
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)

	}
}

func addAgendaItem(w http.ResponseWriter, r *http.Request) {
	var bodyValues addItemStruct
	err := json.NewDecoder(r.Body).Decode(&bodyValues)
	if decoderError(w, err) {
		return
	}

	_, err = db.Query("INSERT INTO agenda_items (name, information, due_date) VALUES (?, ?, ?)", bodyValues.Name, bodyValues.Information, bodyValues.Date)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

func updateAgenda(w http.ResponseWriter, r *http.Request) {
	var bodyValues itemStruct
	err := json.NewDecoder(r.Body).Decode(&bodyValues)
	if decoderError(w, err) {
		return
	}

	_, err = db.Query("UPDATE agenda_items SET name = ?, information = ?, due_date = ?, done = ? WHERE id = ?", bodyValues.Name, bodyValues.Information, bodyValues.DueDate, bodyValues.Done, bodyValues.ID)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

func deleteAgendaItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Get("id") != "" {
		_, err := db.Query("DELETE FROM agenda_items WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		w.WriteHeader(204)
	} else {
		w.WriteHeader(400)
	}
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
