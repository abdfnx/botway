<div align="center">
  <h1>botway-rs</h1>
	<p>
		Rust client package for Botway
	</p>
	<br />
	<p>
		<img alt="Crates.io" src="https://img.shields.io/crates/v/botway-rs?logo=rust&style=flat-square">
	</p>
</div>

## Installation

- `Cargo.toml`

```toml
[dependencies]
botway-rs = "0.2.2"
```

- `CLI`

```
cargo add botway-rs
```

## Usage

> after creating a new rust botway project, you need to use your tokens to connect with your bot.

```rust
use teloxide::{prelude::*, utils::command::BotCommands};
use std::error::Error;
use botway_rs::get;

#[tokio::main]
async fn main() {
    pretty_env_logger::init();
    log::info!("Starting command bot...");

    let bot = Bot::new(get("token"));

    teloxide::commands_repl(bot, answer, Command::ty()).await;
}
...
```
