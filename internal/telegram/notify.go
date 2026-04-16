package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type message struct {
	Chat_id string `json:"chat_id"`
	Text    string `json:"text"`
}

func SendAlert(chatID string, token string, msg string) (error error) {

	body := message{
		Chat_id: chatID,
		Text:    msg,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sedMessage", token)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
