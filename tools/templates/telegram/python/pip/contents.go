package pip

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV TELEGRAM_BOT_NAME="%s"
ARG TELEGRAM_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python3-dev libffi-dev git

RUN botway init --docker --name ${TELEGRAM_BOT_NAME}
RUN pip3 install -r requirements.txt

# Add packages you want
# RUN apk add PACKAGE_NAME

EXPOSE 8000

ENTRYPOINT ["python3", "./src/main.py"]`, botName)
}

func MainPyContent() string {
	return `"""
Simple Bot to reply to Telegram messages.
First, a few handler functions are defined. Then, those functions are passed to
the Dispatcher and registered at their respective places.
Then, the bot is started and runs until we press Ctrl-C on the command line.
Usage:
Basic Echobot example, repeats messages.
Press Ctrl-C on the command line or send a signal to the process to stop the
bot.
"""

import logging
import botway

from telegram import Update, ForceReply
from telegram.ext import Updater, CommandHandler, MessageHandler, Filters, CallbackContext

logging.basicConfig(
	format='%(asctime)s - %(name)s - %(levelname)s - %(message)s', level=logging.INFO
)

logger = logging.getLogger(__name__)

# Define a few command handlers. These usually take the two arguments update and
# context.
def start(update: Update, context: CallbackContext) -> None:
	"""Send a message when the command /start is issued."""
	user = update.effective_user
	update.message.reply_markdown_v2(
		fr'Hi {user.mention_markdown_v2()}\!',
		reply_markup=ForceReply(selective=True),
	)

def help_command(update: Update, context: CallbackContext) -> None:
	"""Send a message when the command /help is issued."""
	update.message.reply_text('Help!')

def echo(update: Update, context: CallbackContext) -> None:
	"""Echo the user message."""
	update.message.reply_text(update.message.text)

def main() -> None:
	"""Start the bot."""
	# Create the Updater and pass it your bot's token.
	updater = Updater(botway.GetToken())

	# Get the dispatcher to register handlers
	dispatcher = updater.dispatcher

	# on different commands - answer in Telegram
	dispatcher.add_handler(CommandHandler("start", start))
	dispatcher.add_handler(CommandHandler("help", help_command))

	# on non command i.e message - echo the message on Telegram
	dispatcher.add_handler(MessageHandler(Filters.text & ~Filters.command, echo))

	# Start the Bot
	updater.start_polling()

	# Run the bot until you press Ctrl-C or the process receives SIGINT,
	# SIGTERM or SIGABRT. This should be used most of the time, since
	# start_polling() is non-blocking and will stop the bot gracefully.
	updater.idle()

if __name__ == '__main__':
	main()`
}

func Resources() string {
	return `# Botway Telegtam (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Telegram Bot API for Python](https://github.com/python-telegram-bot/python-telegram-bot)
- [Python-Telegram-Bot Website](https://python-telegram-bot.org)
- [Telegram Group](https://telegram.me/pythontelegrambotgroup)

## Examples

- [A collection of examples written with Python-Telegram-Bot](https://github.com/python-telegram-bot/python-telegram-bot/tree/master/examples)

big thanks to [**@python-telegram-bot**](https://github.com/python-telegram-bot) org`
}
