package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Value string
}

type Result struct {
	Message string
}

func sendRequest(request Request, url string) (*Result, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result := Result{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
