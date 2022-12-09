# Botwayrb

> Ruby client package for Botway.

## Dependencies

- Ruby >= 2.6 supported
- An installed build system for native extensions (on Windows, make sure you download the "Ruby+Devkit" version of [RubyInstaller](https://rubyinstaller.org/downloads/))

## Installation

### With Bundler

Using [Bundler](https://bundler.io/#getting-started), you can add bwrb to your Gemfile:

    gem "bwrb"

And then install via `bundle install`.

### With Gem

Alternatively, while Bundler is the recommended option, you can also install bwrb without it.

#### Linux / macOS

    gem install bwrb

#### Windows

> **Make sure you have the DevKit installed!**

    gem install bwrb --platform=ruby

```ruby
gem "bwrb"
```

## Usage

> this is an example of botway discord ruby template

```ruby
require "discordrb"
require "bwrb"

bw = Botwayrb::Core.new
bot = Discordrb::Bot.new token: bw.get_token

bot.message(with_text: "ping") do |event|
  event.respond "pong!"
end

bot.run
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
