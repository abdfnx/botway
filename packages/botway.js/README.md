<div align="center">
  <h1>botway.js</h1>
	<p>
		botway nodejs client library
	</p>
	<br />
	<p>
		<img alt="npm" src="https://img.shields.io/npm/v/botway.js?logo=npm&style=flat-square">
	</p>
</div>

```bash
# npm
npm i botway.js

# yarn
yarn add botway.js

# pnpm
pnpm add botway.js
```

## Usage

> after creating a new nodejs botway project, you need to use your tokens to connect with your bot.

```js
...
const { GetToken, GetAppId, GetGuildId } = require("botway.js");
const { Client, Intents } = require("discord.js");

const rest = new REST({ version: "10" }).setToken(GetToken());

rest
  // how to get the id of my server? the answer in this discussion https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737
  .put(Routes.applicationGuildCommands(GetAppId(), GetGuildId("YOUR_SERVER_NAME"))
...
```
