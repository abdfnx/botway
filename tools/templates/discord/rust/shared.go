package rust

import "fmt"

func MainRsContent() string {
	return `use serenity::async_trait;
use serenity::model::channel::Message;
use serenity::model::gateway::Ready;
use serenity::prelude::*;

use botway_rs::get_token;

struct Handler;

#[async_trait]
impl EventHandler for Handler {
	// Set a handler for the message event - so that whenever a new message
	// is received - the closure (or function) passed will be called.
	//
	// Event handlers are dispatched through a threadpool, and so multiple
	// events can be dispatched simultaneously.
	async fn message(&self, ctx: Context, msg: Message) {
		if msg.content == "ping" {
			// Sending a message can fail, due to a network error, an
			// authentication error, or lack of permissions to post in the
			// channel, so log to stdout when some error happens, with a
			// description of it.
			if let Err(why) = msg.channel_id.say(&ctx.http, "Pong!").await {
				println!("Error sending message: {:?}", why);
			}
		}
	}

	// Set a handler to be called on the ready event. This is called when a
	// shard is booted, and a READY payload is sent by Discord. This payload
	// contains data like the current user's guild Ids, current user data,
	// private channels, and more.
	//
	// In this case, just print what the current user's username is.
	async fn ready(&self, _: Context, ready: Ready) {
		println!("{} is connected!", ready.user.name);
	}
}

#[tokio::main]
async fn main() {
	// Configure the client with your Discord bot token in the environment.
	let token = get_token;
	// Set gateway intents, which decides what events the bot will be notified about
	let intents = GatewayIntents::GUILD_MESSAGES
		| GatewayIntents::DIRECT_MESSAGES
		| GatewayIntents::MESSAGE_CONTENT;

	// Create a new instance of the Client, logging in as a bot. This will
	// automatically prepend your bot token with "Bot ", which is a requirement
	// by Discord for bot users.
	let mut client =
		Client::builder(&token, intents).event_handler(Handler).await.expect("Err creating client");

	// Finally, start a single shard, and start listening to events.
	//
	// Shards will automatically attempt to reconnect, and will perform
	// exponential backoff until it reconnects.
	if let Err(why) = client.start().await {
		println!("Client error: {:?}", why);
	}
}`
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
