const TelegramBot = require("node-telegram-bot-api");
const botway = require("botway.js");
const request = require("request");

const options = {
  polling: true,
};

const bot = new TelegramBot(botway.GetToken(), options);

// Matches /photo
bot.onText(/\/photo/, function onPhotoText(msg) {
  // From file path
  const photo = request("https://raw.githubusercontent.com/abdfnx/botway/main/templates/telegram/nodejs/assets/bot.gif");

  bot.sendPhoto(msg.chat.id, photo, {
    caption: "I'm a bot!",
  });
});

// Matches /crypto
bot.onText(/\/crypto/, function onFavoriteCryptoText(msg) {
  const opts = {
    reply_to_message_id: msg.message_id,
    reply_markup: JSON.stringify({
      keyboard: [
        ["BTC", "ETH", "LTC"],
        ["EOS", "XRP", "SOL"],
      ],
    }),
  };
  bot.sendMessage(msg.chat.id, "What is your favorite crypto currency?", opts);
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
});
