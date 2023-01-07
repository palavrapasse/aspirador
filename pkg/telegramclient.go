package pkg

import (
	"log"
)

const messageType = "application/json"

type TelegramClient struct {
	loggers []*log.Logger
}

func NewTelegramClient(botToken, chatId string) TelegramClient {
	tw := TelegramWriter{
		botToken: botToken,
		chatId:   chatId,
	}

	loggers := make([]*log.Logger, len(levelPrefix))

	for i, v := range levelPrefix {
		loggers[i] = log.New(tw, v, defaultLoggerFlag)
	}

	return TelegramClient{
		loggers: loggers,
	}
}

func (tc TelegramClient) Write(ar Record) {
	tc.loggers[ar.Level].Println(ar.Message)
}
