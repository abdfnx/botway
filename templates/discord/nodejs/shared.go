package nodejs

import "github.com/abdfnx/botway/templates"

var Packages = "discord.js @discordjs/rest @discordjs/builders discord-api-types discord-rpc zlib-sync erlpack bufferutil utf-8-validate @discordjs/voice libsodium-wrappers @discordjs/opus sodium botway.js"

func IndexJSContent() string {
	return templates.Content("discord", "nodejs", "src/index.js", "")
}

func Resources() string {
	return `# Botway Discord (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [A powerful JavaScript library for interacting with the Discord API](https://github.com/discordjs/discord.js)
- [Discord.js Website](https://discord.js.org)
- [Discord.js Docs](https://discord.js.org/#/docs)
- [Discord.js Guide](https://discordjs.guide)
- [Discord Server](https://discord.gg/djs)

big thanks to [**@discordjs**](https://github.com/discordjs) org`
}
