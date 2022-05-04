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

```toml
[dependencies]
botway-rs = "0.1"
```

## Usage

> after creating a new rust botway project, you need to use your tokens to connect with your bot.

```rust
use futures::StreamExt;
use telegram_bot::*;

use botway_rs::get;

#[tokio::main]
async fn main() -> Result<(), Error> {
	let api = Api::new(get("token"));
	let mut stream = api.stream();
	...
}
...
```
