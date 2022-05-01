package python

import "github.com/abdfnx/botway/tools/templates"

func MainPyContent() string {
	return templates.Content("discord", "python", "src/main.py", "")
}

func Resources() string {
	return `# Botway Discord (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [An API wrapper for Discord written in Python.](https://github.com/Rapptz/discord.py)
- [Discord.py Website](https://discordpy.rtfd.org/en/latest)
- [Discord Server](https://discord.gg/r3sSKJJ)

## Examples

- [A collection of example programs written with Discord.py](https://github.com/Rapptz/discord.py/tree/master/examples)

big thanks to [**@Rapptz**](https://github.com/Rapptz)`
}
