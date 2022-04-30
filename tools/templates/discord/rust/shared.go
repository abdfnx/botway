package rust

import "github.com/abdfnx/botway/tools/templates"

func MainRsContent() string {
	return templates.Content("discord", "rust", "src/main.rs", "")
}

func CargoFileContent(botName string) string {
	return templates.Content("discord", "rust", "Cargo.toml", botName)
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
