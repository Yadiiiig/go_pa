package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	token         = "3truutQfPlTnVVHkm9uo2HET"
	authToken     = "xoxb-1809789588064-1785984541650-TlX0AWP0bOkcuIW9s6FOM2YQ"
	authorization = "Willem"
	url           = "http://127.0.0.1:8000/"
	webhookURL    = "https://hooks.slack.com/services/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"
	apiURL        = "https://slack.com/api/"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", slackHandlerSlash).Methods("POST")
	router.HandleFunc("/interactive", slackHandlerSubmissions).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func slackHandlerSlash(w http.ResponseWriter, r *http.Request) {
	command := r.FormValue("command")

	switch command {
	case "/print":
		fmt.Fprint(w, "You typed: \n"+r.FormValue("text"))
	case "/agenda_items_between":
		getAgendaItems(r.FormValue("trigger_id"))
	case "/create_agenda_item":
		getCreateAgendaItem(r.FormValue("trigger_id"))
	default:
		w.WriteHeader(404)
	}
}

func slackHandlerSubmissions(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("payload") != "" {
		payload := []byte(r.FormValue("payload"))
		var raw map[string]interface{}

		if err := json.Unmarshal(payload, &raw); err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
		}

		if raw["type"] != "view_submission" {
			w.WriteHeader(200)
			return
		}

		payloadStruct := modalPayload{}
		tempBytes, _ := json.Marshal(raw["view"])
		json.Unmarshal(tempBytes, &payloadStruct)
		id := payloadStruct.CallbackID
		switch id {
		case "get_agenda_item_between":
			getAgendaBetween(payloadStruct.State, w)
			w.WriteHeader(200)

		case "add_agenda_item":
			addAgendaItem(payloadStruct.State, w)
			w.WriteHeader(200)
		}

	} else {
		w.WriteHeader(404)
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
