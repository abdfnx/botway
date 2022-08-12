<div align="center">
  <h1>Botnim</h1>
	<p>
		Nim client package for Botway
	</p>
</div>

## Usage

> after creating a new nim botway project, you need to use your tokens to connect with your bot.

```nim
import dimscord, asyncdispatch, strutils, options
import botnim

let discord = newDiscordClient(GetToken())
...
```
