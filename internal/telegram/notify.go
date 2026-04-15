package telegram

type Message struct {
	chat_id string
	text    string
}

func SendAlert(chatID string, token string, message string) (error error) {

	return nil
}
