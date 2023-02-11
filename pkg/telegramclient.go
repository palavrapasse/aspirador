package pkg

const messageType = "application/json"

type TelegramClient struct {
	loggers LevelLogger
}

func NewTelegramClient(botToken, chatId string, levels ...Level) TelegramClient {
	tw := TelegramWriter{
		botToken: botToken,
		chatId:   chatId,
	}

	loggers := NewLevelLogger(tw, defaultLoggerFlag, levels)

	return TelegramClient{
		loggers: loggers,
	}
}

func (tc TelegramClient) Write(ar Record) {
	logger, exists := tc.loggers[ar.Level]

	if !exists {
		return
	}

	logger.Println(ar.Message)
}
