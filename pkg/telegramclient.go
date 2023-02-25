package pkg

const messageType = "application/json"

type TelegramClient struct {
	loggers       LevelLogger
	patternLayout PatternLayout
}

func NewTelegramClient(botToken, chatId string, levels ...Level) TelegramClient {
	tw := TelegramWriter{
		botToken: botToken,
		chatId:   chatId,
	}

	loggers := NewLevelLogger(tw, defaultLoggerFlag, levels)

	return TelegramClient{
		loggers:       loggers,
		patternLayout: defaultPatternLayout,
	}
}

func (tc *TelegramClient) SetPatternLayout(p PatternLayout) {
	tc.patternLayout = p
}

func (tc TelegramClient) Write(ar Record) {
	message := tc.patternLayout.FormatRecord(ar)

	tc.loggers[ar.Level].Println(message)
}

func (tc TelegramClient) SupportsLevel(l Level) bool {
	return tc.loggers.ContainsLevel(l)
}
