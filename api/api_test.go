package main

import (
	"net/http"
	"testing"
)

func TestAuthentication(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/get_notes", nil)
	//req.Header.Set("Authorization", "Willem")
	res, _ := client.Do(req)
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("Expected %d, received %d", 403, res.StatusCode)
	}
}
