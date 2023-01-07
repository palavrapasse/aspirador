package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const messageType = "application/json"

type TelegramClient struct {
	url    string
	chatId string
	logs   []string
}

func NewTelegramClient(botToken, chatId string) TelegramClient {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	return TelegramClient{
		url:    url,
		logs:   levelPrefix,
		chatId: chatId,
	}
}

func (tc TelegramClient) Write(ar Record) {
	log := tc.logs[ar.Level]

	body, err := json.Marshal(map[string]string{
		"chat_id": tc.chatId,
		"text":    fmt.Sprintf("%s%s", log, ar.Message),
	})

	if err != nil {
		return
	}

	response, err := http.Post(tc.url, messageType, bytes.NewBuffer(body))

	if err != nil {
		return
	}

	defer response.Body.Close()

}
