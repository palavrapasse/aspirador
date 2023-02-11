# aspirador

Shared logging module

## Hooks

This repository is configured with client-side Git hooks which you need to install by running the following command:

```bash
./hooks/INSTALL
```

## Features

- Log to console
- Log to file
- Log to Telegram chat
- Enumerate the log's levels 


## Usage

First, include `aspirador` module.

```bash
go get github.com/palavrapasse/aspirador
```

Then import it to your project.

```golang
import aspirador "github.com/palavrapasse/aspirador/pkg"
```


## Examples

Log to console, file and Telegram chat.

```golang
package main

import (
	aspirador "github.com/palavrapasse/aspirador/pkg"
)

func main() {
	fileClient, err := aspirador.NewFileClient("filename.log") // Will print all Levels (TRACE, INFO, WARNING, ERROR) logs
	if err != nil {
		return
	}

	consoleClient := aspirador.NewConsoleClient(aspirador.WARNING, aspirador.ERROR) // Will only print Warning and Error logs

	telegramClient := aspirador.NewTelegramClient("telegram_bot_token", "chat_id", aspirador.ERROR) // Will only print Error logs

	clients := []aspirador.Client{fileClient, consoleClient, telegramClient}

	as := aspirador.WithClients(clients)
	as.Trace("Trace message")
	as.Info("info message")
	as.Warning("Warning message")
	as.Error("Error message")
}
```