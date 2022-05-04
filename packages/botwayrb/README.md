# Botwayrb

> Ruby client package for Botway.

## Dependencies

* Ruby >= 2.6 supported
* An installed build system for native extensions (on Windows, make sure you download the "Ruby+Devkit" version of [RubyInstaller](https://rubyinstaller.org/downloads/))

## Installation

### With Bundler

Using [Bundler](https://bundler.io/#getting-started), you can add botwayrb to your Gemfile:

    gem "botwayrb"

And then install via `bundle install`.

### With Gem

Alternatively, while Bundler is the recommended option, you can also install botwayrb without it.

#### Linux / macOS

    gem install botwayrb

#### Windows

> **Make sure you have the DevKit installed! See the [Dependencies](https://github.com/shardlab/botwayrb#dependencies) section)**
    gem install botwayrb --platform=ruby

```ruby
gem "botwayrb"
```

## Usage

> this is an example of botway discord ruby template

```ruby
require "discordrb"
require "botwayrb"

bw = Botwayrb::Core.new
bot = Discordrb::Bot.new token: bw.get_token

bot.message(with_text: "ping") do |event|
  event.respond "pong!"
end

bot.run
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
