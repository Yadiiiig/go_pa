package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	token         = ""
	authToken     = ""
	authorization = "Willem"
	url           = "http://127.0.0.1:8000/"
	webhookURL    = ""
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
			w.WriteHeader(200)
		}

		//fmt.Println(payloadStruct.State)

		// switch {
		// case action == "submit_agenda_between":
		// 	getAgendaBetween(raw["state"], w)
		// default:
		// 	w.WriteHeader(404)
		// }
		// if r.FormValue("view") != "" {
		// 	payload := []byte(r.FormValue("view"))
		// 	payloadView := []modalReturnStruct{}

		// 	if err := json.Unmarshal(payload, &payloadView); err != nil {
		// 		fmt.Println(err)
		// 		w.WriteHeader(404)
		// 	}

		// 	fmt.Println(payloadView[0].CallbackID)
		// 	w.WriteHeader(200)
		// } else {
		// 	w.WriteHeader(404)
		// }
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
