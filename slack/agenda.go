package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAgendaItems(triggerID string) {
	err := sendSlackNotificationModal(webhookURL, getAgendaBetweenModal, triggerID)
	if err != nil {
		fmt.Println(err)
	}
}

func getAgendaBetween(payload interface{}, w http.ResponseWriter) {
	bodyValues := []itemStruct{}

	payloadJSON := agendaBetweenPayload{}
	test, _ := json.Marshal(payload)

	if err := json.Unmarshal(test, &payloadJSON); err != nil {
		fmt.Println(err)
	}

	after := payloadJSON.Values.Dates.After.SelectedDate
	before := payloadJSON.Values.Dates.Before.SelectedDate
	fmt.Println(after, before)
	type values map[string]interface{}
	tempValues := values{"after": after, "before": before}

	response := getRequest("get_agenda_items", tempValues)

	if err := json.NewDecoder(response.Body).Decode(&bodyValues); err != nil {
		fmt.Println(err)
	}

	returnMessage := ""

	for i := 0; i < len(bodyValues); i++ {

		temp := fmt.Sprintf("ID: %v - Title: %s \n Content: %s \n Due date: %s - Done: %t \n\n ", bodyValues[i].ID, bodyValues[i].Name, bodyValues[i].Information, bodyValues[i].DueDate, bodyValues[i].Done)
		returnMessage += temp
	}

	sendSlackNotificationText(webhookURL, returnMessage)
}

func getCreateAgendaItem(triggerID string) {
	err := sendSlackNotificationModal(webhookURL, addAgendaItemModal, triggerID)
	if err != nil {
		fmt.Println(err)
	}
}

type itemStruct struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Information string `db:"information"`
	DueDate     string `db:"due_date"`
	Done        bool   `db:"done"`
}
