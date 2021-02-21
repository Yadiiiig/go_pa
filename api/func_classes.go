package main

import (
	"encoding/json"
	"net/http"
)

func getClasses(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	selectedItems := []classesStruct{}

	if query.Get("id") != "" {
		err := db.Select(&selectedItems, "SELECT * FROM classes WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)
	} else {
		err := db.Select(&selectedItems, "SELECT * FROM classes")
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)
	}

	// Could still add a search on name since they would be unique (normally)

}

func updateClass(w http.ResponseWriter, r *http.Request) {
	var bodyValues classesStruct
	err := json.NewDecoder(r.Body).Decode(&bodyValues)
	if decoderError(w, err) {
		return
	}

	_, err = db.Query("UPDATE classes SET name = ?, teacher = ? WHERE id = ?", bodyValues.Name, bodyValues.Teacher, bodyValues.ID)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
}

type classesStruct struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Teacher string `db:"teacher" json:"teacher"`
}
