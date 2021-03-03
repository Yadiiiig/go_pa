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
