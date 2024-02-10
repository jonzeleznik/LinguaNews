package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

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

func ChatGpt(article string) (respone, error) {
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
				Content: "You are an assistant that rewrites English articles and translates them into Slovenian",
			},
			{
				Role:    "user",
				Content: article,
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
	req.Header.Set("Authorization", "Bearer "+"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")
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
