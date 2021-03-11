package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type slackRequestBody struct {
	Blocks string `json:"blocks"`
}

type slackRequestBodyText struct {
	Text string `json:"text"`
}

type slackRequestBodyModal struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Submit string `json:"submit"`
	Close  string `json:"close"`
	Blocks string `json:"blocks"`
}

type testing struct {
	TriggerID string `json:"trigger_id"`
	View      string `json:"view"`
}

type fullBody struct {
	Token   string `json:"token"`
	Payload testing
}

func sendSlackNotification(webhookURL string, block string) error {
	slackBody, _ := json.Marshal(slackRequestBody{Blocks: block})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
	if buf.String() != "ok" {
		fmt.Println("Slack said no")
	}
	return nil
}

func sendSlackNotificationText(webhookURL string, text string) error {
	slackBody, _ := json.Marshal(slackRequestBodyText{Text: text})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
	if buf.String() != "ok" {
		fmt.Println("Slack said no")
	}
	return nil
}

func sendSlackNotificationModal(webhookURL string, block string, triggerID string) error {
	slackBody, _ := json.Marshal(testing{
		TriggerID: triggerID,
		View:      block,
	})

	req, err := http.NewRequest(http.MethodPost, apiURL+"views.open", bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{Timeout: 10 * time.Second}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(resp.Body)
	// fmt.Println(buf.String())
	// if buf.String() != "ok" {
	// 	fmt.Println("Slack said no")
	// }
	return nil
}
