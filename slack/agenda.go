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

func getCreateAgendaItem(triggerID string) {
	err := sendSlackNotificationModal(webhookURL, addAgendaItemModal, triggerID)
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

func addAgendaItem(payload interface{}, w http.ResponseWriter) {
	payloadJSON := addAgendaItemPayload{}
	test, _ := json.Marshal(payload)

	if err := json.Unmarshal(test, &payloadJSON); err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	postBody, _ := json.Marshal(addAgendaStruct{
		Title:       payloadJSON.Values.Title.TitleInput.Value,
		Information: payloadJSON.Values.Information.InformationInput.Value,
		DueDate:     payloadJSON.Values.Date.DueDateInput.SelectedDate,
	})

	_ = postRequest("add_agenda_items", postBody)
}

type itemStruct struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Information string `db:"information"`
	DueDate     string `db:"due_date"`
	Done        bool   `db:"done"`
}

type addAgendaStruct struct {
	Title       string `json:"name"`
	Information string `json:"info"`
	DueDate     string `json:"date"`
}
