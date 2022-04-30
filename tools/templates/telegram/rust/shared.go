package rust

import "fmt"

func MainRsContent() string {
	return `use futures::StreamExt;
use telegram_bot::*;

use botway_rs::get_token;

#[tokio::main]
async fn main() -> Result<(), Error> {
	let api = Api::new(get_token());

	// Fetch new updates via long poll method
	let mut stream = api.stream();

	while let Some(update) = stream.next().await {
		// If the received update contains a new message...
		let update = update?;

		if let UpdateKind::Message(message) = update.kind {
			if let MessageKind::Text { ref data, .. } = message.kind {
				// Print received text message to stdout.
				println!("<{}>: {}", &message.from.first_name, data);

				// Answer message with "Hi".
				api.send(message.text_reply(format!(
					"Hi, {}! You just wrote '{}'",
					&message.from.first_name, data
				)))
				.await?;
			}
		}
	}

	Ok(())
}`
}

func CargoFileContent(botName string) string {
	return fmt.Sprintf(`[package]
name = "%s"
version = "0.1.0"
edition = "2021"

[dependencies]
telegram-bot = "0.7"
futures = "0.3.21"
botway-rs = "0.0.1"`, botName)
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
