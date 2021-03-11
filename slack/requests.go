package main

import (
	"fmt"
	"net/http"
)

func getRequest(path string, params map[string]interface{}) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("authorization", authorization)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}
