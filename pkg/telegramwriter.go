package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramWriter struct {
	botToken string
	chatId   string
}

func (tw TelegramWriter) Write(p []byte) (n int, err error) {

	body, err := json.Marshal(map[string]string{
		"chat_id": tw.chatId,
		"text":    string(p[:]),
	})

	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tw.botToken)

	go asyncPost(url, body)

	return len(p), nil
}

func asyncPost(url string, body []byte) {
	response, err := http.Post(url, messageType, bytes.NewBuffer(body))

	if err != nil {
		return
	}

	if response.Body != nil {
		response.Body.Close()
	}

}
