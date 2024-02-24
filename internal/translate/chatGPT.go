package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const api_token = ""

type payload struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type respone struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Finish_reason string `json:"finish_reason"`
	} `json:"choices"`
}

func ChatGpt(title, description, content string) (respone, error) {
	var body respone
	client := &http.Client{}

	data := payload{
		Model: "gpt-3.5-turbo-0125",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: "You are an assistant that rewrites English articles and translates them into Slovenian. The articles are about movie industry and new releases.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Title: %s \nDescription: %s \nContent: %s", title, description, content),
			},
			{
				Role:    "user",
				Content: "Translate further.",
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return body, err
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(jsonData))

	if err != nil {
		return body, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_token)
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	err = json.Unmarshal([]byte(bodyText), &body)
	if err != nil {
		return body, err
	}

	return body, nil
}