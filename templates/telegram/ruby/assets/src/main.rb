# frozen_string_literal: true

require "telegram-bot-ruby"
require "botwayrb"

bw = Botwayrb::Core.new

Telegram::Bot::Client.run(bw.get_token) do |bot|
  bot.listen do |message|
    case message.text
    when "/start"
      bot.api.send_message(chat_id: message.chat.id, text: "Hello, #{message.from.first_name}")
    when "/stop"
      bot.api.send_message(chat_id: message.chat.id, text: "Bye, #{message.from.first_name}")
    end
  end
end
