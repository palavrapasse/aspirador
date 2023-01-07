package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const tokenEnvKey = "telegram_token"
const messageType = "application/json"

type TelegramClient struct {
	url    string
	chatId string
	logs   []string
}

func NewTelegramClient(chatId string) TelegramClient {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv(tokenEnvKey))

	return TelegramClient{
		url:    url,
		logs:   levelPrefix,
		chatId: chatId,
	}
}

func (tw TelegramClient) Write(ar Record) {
	log := tw.logs[ar.Level]

	body, err := json.Marshal(map[string]string{
		"chat_id": tw.chatId,
		"text":    log + ar.Message,
	})

	if err != nil {
		return
	}

	response, err := http.Post(tw.url, messageType, bytes.NewBuffer(body))

	if err != nil {
		return
	}

	defer response.Body.Close()

}
