package pip

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python3-dev libffi-dev git

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN pip3 install -r requirements.txt

# Add packages you want
# RUN apk add PACKAGE_NAME

EXPOSE 8000

ENTRYPOINT ["python3", "./src/main.py"]`, botName)
}

func MainPyContent() string {
	return `import discord
import botway

intents = discord.Intents.default()

client = discord.Client(intents=intents)

@client.event
async def on_ready():
	print(f'We have logged in as {client.user}')

@client.event
async def on_message(message):
	if message.author == client.user:
		return

	if message.content.startswith('Hi'):
		await message.channel.send('Hello!')

client.run(botway.GetToken())`
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

big thanks to [**@Rapptz**](https://github.com/Rapptz) org`
}
