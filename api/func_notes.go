package main

import (
	"encoding/json"
	"net/http"
)

func getNotes(w http.ResponseWriter, r *http.Request) {
	selectedItems := []noteStruct{}
	query := r.URL.Query()

	if query.Get("id") != "" {
		err := db.Select(&selectedItems, "SELECT * FROM note_items WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)

	} else if query.Get("disabled") != "" {
		err := db.Select(&selectedItems, "SELECT * FROM note_items WHERE disabled = 1")
		if databaseErrorRequest(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)
	} else {
		err := db.Select(&selectedItems, "SELECT * FROM note_items")
		if databaseErrorRequest(w, err) {
			return
		}

		checkEmpty(w, len(selectedItems))
		json.NewEncoder(w).Encode(selectedItems)
	}
}

// Could add the disabled state, but I see no usecase in the near future
func addNote(w http.ResponseWriter, r *http.Request) {
	var bodyValues addNoteStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	_, err := db.Query("INSERT INTO note_items (title, content) VALUES (?, ?)", bodyValues.Title, bodyValues.Content)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	var bodyValues noteStruct
	json.NewDecoder(r.Body).Decode(&bodyValues)

	_, err := db.Query("UPDATE note_items SET title = ?, content = ?, disabled = ? WHERE id = ?", bodyValues.Title, bodyValues.Content, bodyValues.Disabled, bodyValues.ID)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Technicly not needed since usage of API will always make sure to have one
	if query.Get("id") == "" {
		w.WriteHeader(400)
		return
	}
	_, err := db.Query("DELETE FROM note_items WHERE id = ?", query.Get("id"))
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

type noteStruct struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Content  string `db:"content" json:"content"`
	Disabled int    `db:"disabled" json:"disabled"`
}

type addNoteStruct struct {
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}
