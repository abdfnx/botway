package nodejs

import "fmt"

var Packages = "node-telegram-bot-api botway.js"

func IndexJSContent() string {
	return fmt.Sprintf(`const TelegramBot = require("node-telegram-bot-api");
const botway = require("botway.js")
const request = require("request");

const options = {
	polling: true,
};

const bot = new TelegramBot(botway.GetToken(), options);

// Matches /photo
bot.onText(/\/photo/, function onPhotoText(msg) {
	// From file path
	const photo = %s;

	bot.sendPhoto(msg.chat.id, photo, {
		caption: "I'm a bot!",
	});
});

// Matches /audio
bot.onText(/\/audio/, function onAudioText(msg) {
	// From HTTP request
	const url = "https://upload.wikimedia.org/wikipedia/commons/c/c8/Example.ogg";
	const audio = request(url);

	bot.sendAudio(msg.chat.id, audio);
});

// Matches /echo [whatever]
bot.onText(/\/echo (.+)/, function onEchoText(msg, match) {
	const resp = match[1];

	bot.sendMessage(msg.chat.id, resp);
});

// Matches /editable
bot.onText(/\/editable/, function onEditableText(msg) {
	const opts = {
		reply_markup: {
			inline_keyboard: [
				[
					{
						text: "Edit Text",
						// we shall check for this value when we listen
						// for "callback_query"
						callback_data: "edit",
					},
				],
			],
		},
	};

	bot.sendMessage(msg.from.id, "Original Text", opts);
});

// Handle callback queries
bot.on("callback_query", function onCallbackQuery(callbackQuery) {
	const action = callbackQuery.data;
	const msg = callbackQuery.message;
	const opts = {
		chat_id: msg.chat.id,
		message_id: msg.message_id,
	};

	let text;

	if (action === "edit") {
		text = "Edited Text";
	}

	bot.editMessageText(text, opts);
});`, "`${__dirname}/assets/photo.gif`")
}

func Resources() string {
	return `# Botway Telegram (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Telegram Bot API for NodeJS](https://github.com/https://github.com/yagop/node-telegram-bot-api)
- [node-telegram-bot-api Docs](https://github.com/yagop/node-telegram-bot-api/tree/master/doc)
- [node-telegram-bot-api Help Information](https://github.com/yagop/node-telegram-bot-api/blob/master/doc/help.md)
- [Tutorials](https://github.com/yagop/node-telegram-bot-api/tree/master/doc/tutorials.md)
- [node-telegram-bot-api Telegram Channel](https://t.me/node_telegram_bot_api)
- [node-telegram-bot-api Telegram Group](https://t.me/ntbasupport)

## Examples

[Examples](https://github.com/yagop/node-telegram-bot-api/tree/master/examples)

big thanks to [**@yagop**](https://github.com/yagop)`
}
