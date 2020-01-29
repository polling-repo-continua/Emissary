package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Telegram Send messages via telegram
func Telegram(chatID string, apiKey string, message string) (*http.Response, error) {

	jayson := map[string]interface{}{
		"chat_id": chatID,
		"text":    message,
	}
	js, _ := json.Marshal(jayson)
	endpoint := "https://api.telegram.org/bot" + apiKey + "/sendMessage"

	return request(endpoint, string(js))
}

// Slack Send messages via Slack
func Slack(message string, webhook string) (*http.Response, error) {
	js := `{"text":"` + message + `"}`
	return request(webhook, js)
}

func request(endpoint string, data string) (*http.Response, error) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var resp *http.Response
	var err error

	client := &http.Client{Transport: tr}

	resp, err = client.Post(endpoint, "application/json", strings.NewReader(data))

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	return resp, nil
}