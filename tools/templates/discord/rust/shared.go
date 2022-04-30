package rust

import "fmt"

func MainRsContent() string {
	return ``
}

func CargoFileContent(botName string) string {
	return fmt.Sprintf(`[package]
name = "%s"
version = "0.1.0"
edition = "2021"

[dependencies]
serenity = { version = "0.11", default-features = false, features = ["client", "gateway", "rustls_backend", "model", "voice"] }
tokio = { version = "1.0", features = ["full"] }
botway-rs = "0.0.1"
songbird = "0.2.2"`, botName)
}

func Resources() string {
	return `# Botway Discord (Rust 🦀) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [A Rust library for the Discord API](https://github.com/serenity-rs/serenity)
- [An async Rust library for the Discord voice API](https://github.com/serenity-rs/songbird)
- [Discord Server](https://discord.gg/serenity-rs)

## Examples

- [serenity examples](https://github.com/serenity-rs/serenity/tree/current/examples)
- [songbird examples](https://github.com/serenity-rs/songbird/tree/current/examples)

big thanks to [**@serenity-rs**](https://github.com/serenity-rs) org`
}
