package rust

import "github.com/abdfnx/botway/tools/templates"

func MainRsContent() string {
	return templates.Content("telegram", "rust", "src/main.rs", "")
}

func CargoFileContent(botName string) string {
	return templates.Content("telegram", "rust", "Cargo.toml", botName)
}

func Resources() string {
	return `# Botway Telegram (Rust 🦀) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Rust Library for creating a Telegram Bot](https://github.com/telegram-rs/telegram-bot)

## Examples

- [Telegram-rs examples](https://github.com/telegram-rs/telegram-bot/tree/main/lib/examples)

big thanks to [**@telegram-rs**](https://github.com/telegram-rs) org`
}
