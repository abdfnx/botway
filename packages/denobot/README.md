<div align="center">
  <h1>denobot</h1>
	<p>
		Deno client module for Botway.
	</p>
	<br />
</div>

## Usage

> after creating a new deno botway project, you need to use your tokens to connect with your bot.

```js
...
import { createBot, enableCachePlugin, getToken, getAppId } from "./deps.ts";
import { events } from "./src/events/mod.ts";

...

export const bot = enableCachePlugin(
	createBot({
		token: getToken(),
		botId: getAppId(),
		intents: [],
		events,
	}),
);
...
```
