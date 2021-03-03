package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func getAgendaItems() {
	webhookURL := "<WEBHOOK_URL>"
	err := sendSlackNotification(webhookURL, getAgendaBetweenModal)
	if err != nil {
		log.Fatal(err)
	}
}

func getAgendaBetween(payload interface{}) {
	payloadJSON := agendaBetweenPayload{}
	test, _ := json.Marshal(payload)

	if err := json.Unmarshal(test, &payloadJSON); err != nil {
		fmt.Println(err)
	}

	fmt.Println(payloadJSON.Values.Dates.After.SelectedDate)
	fmt.Println(payloadJSON.Values.Dates.Before.SelectedDate)
}
