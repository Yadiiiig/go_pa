package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	token = "<TOKEN_HERE>"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", slackHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("payload") != "" {
		payload := []byte(r.FormValue("payload"))
		payloadActions := []actionsStruct{}
		var raw map[string]interface{}

		if err := json.Unmarshal(payload, &raw); err != nil {
			fmt.Println(err)
		}

		tempBytes, _ := json.Marshal(raw["actions"])
		json.Unmarshal(tempBytes, &payloadActions)
		action := payloadActions[0].ActionID

		switch {
		case action == "submit_agenda_between":
			getAgendaBetween(raw["state"])
		}

	} else {
		command := r.FormValue("command")

		switch {
		case command == "/print":
			fmt.Fprint(w, "You typed: \n"+r.FormValue("text"))
		case command == "/agenda_items_between":
			getAgendaItems()
		}
	}

}

func authenticationCheck(request http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.FormValue("token")
		if auth != token {
			w.WriteHeader(403)
			return
		}

		request(w, r)
	})
}
