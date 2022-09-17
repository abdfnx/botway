# frozen_string_literal: true

require_relative "bwrb/version"
require "yaml"
require "json"

module Botwayrb
  BotwatConfig = JSON.parse(File.read(File.join(File.expand_path("~"), ".botway", "botway.json")))

  class Error < StandardError; end

  class Core
    def get_bot_info(value)
      data = YAML.load_file(".botway.yaml")

      data["bot"][value]
    end

    def get_token()
      BotwatConfig["botway"]["bots"][get_bot_info("name")]["bot_token"]
    end

    def get_app_id()
      if get_bot_info("type") == "slack"
        BotwatConfig["botway"]["bots"][get_bot_info("name")]["bot_app_token"]
      else
        BotwatConfig["botway"]["bots"][get_bot_info("name")]["bot_app_id"]
      end
    end

    def get_guild_id(serverName)
      if get_bot_info("type") != "discord"
        raise Error, "ERROR: This function/feature is only working with discord bots"
      else
        BotwatConfig["botway"]["bots"][get_bot_info("name")]["guilds"][serverName]["server_id"]
      end
    end

    def get_signing_secret()
      if get_bot_info("type") != "slack"
        raise Error, "ERROR: This function/feature is only working with slack bots"
      else
        BotwatConfig["botway"]["bots"][get_bot_info("name")]["signing_secret"]
      end
    end
  end
end
