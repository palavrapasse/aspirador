package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
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

	resultChan := make(chan *http.Response, 1)

	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return asyncPost(url, body, resultChan) })

	err = errGrp.Wait()
	if err != nil {
		return 0, err
	}

	response := <-resultChan
	defer response.Body.Close()

	return len(p), nil
}

func asyncPost(url string, body []byte, rc chan *http.Response) error {
	response, err := http.Post(url, messageType, bytes.NewBuffer(body))

	if err == nil {
		rc <- response
	}

	return err
}
