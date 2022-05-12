# frozen_string_literal: true

require "discordrb"
require "botwayrb"

bw = Botwayrb::Core.new
bot = Discordrb::Bot.new token: bw.get_token

bot.message(with_text: "ping") do |event|
  event.respond "pong!"
end

bot.run
