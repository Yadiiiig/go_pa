package main

import (
	"encoding/json"
	"net/http"
	"time"
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

func addClass(w http.ResponseWriter, r *http.Request) {
	var bodyValues classesStruct
	err := json.NewDecoder(r.Body).Decode(&bodyValues)
	if decoderError(w, err) {
		return
	}

	_, err = db.Query("INSERT INTO classes (name, teacher) VALUES (?, ?)", bodyValues.Name, bodyValues.Teacher)
	if databaseErrorRequest(w, err) {
		return
	}

	w.WriteHeader(204)
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

func deleteClass(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Get("id") != "" {
		_, err := db.Query("DELETE FROM classes WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		w.WriteHeader(204)
	} else {
		w.WriteHeader(400)
	}
}

func getRoster(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	selectedItems := []rosterStruct{}

	if query.Get("day") != "" {
		err := db.Select(&selectedItems, "SELECT * FROM classes WHERE id = ?", query.Get("id"))
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)
	} else {
		err := db.Select(&selectedItems, "SELECT class_hours.day, class_hours.hour, class_hours.location, classes.name, classes.teacher FROM class_hours INNER JOIN classes ON class_hours.class_id = classes.id")
		if databaseErrorRequest(w, err) {
			return
		}

		if checkEmpty(w, len(selectedItems)) {
			return
		}
		json.NewEncoder(w).Encode(selectedItems)
	}
}

type classesStruct struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Teacher string `db:"teacher" json:"teacher"`
}

type rosterStruct struct {
	Day      int       `db:"day" json:"day"`
	Hour     time.Time `db:"hour" json:"hour"`
	Location string    `db:"location" json:"location"`
	Name     string    `db:"name" json:"name"`
	Teacher  string    `db:"teacher" json:"teacher"`
}
