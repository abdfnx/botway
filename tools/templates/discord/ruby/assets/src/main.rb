# frozen_string_literal: true

require 'discordrb'
require 'botwayrb'

bot = Discordrb::Bot.new token: botwayrb.getToken()

bot.message(with_text: 'ping') do |event|
  event.respond 'pong!'
end

bot.run
