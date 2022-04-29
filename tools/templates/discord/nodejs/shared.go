package js

import "fmt"

var Packages = "discord.js @discordjs/rest @discordjs/builders discord-api-types discord-rpc zlib-sync erlpack bufferutil utf-8-validate @discordjs/voice libsodium-wrappers @discordjs/opus sodium botway.js"

func IndexJSContent() string {
	return fmt.Sprintf(`const { SlashCommandBuilder } = require("@discordjs/builders");
const { REST } = require("@discordjs/rest");
const { Routes } = require("discord-api-types/v9");
const { GetToken, GetClientId, GetGuildId } = require("botway.js");
const { Client, Intents } = require("discord.js");

const commands = [
	new SlashCommandBuilder()
	.setName("ping")
	.setDescription("Replies with pong!"),
	new SlashCommandBuilder()
	.setName("server")
	.setDescription("Replies with server info!"),
].map((command) => command.toJSON());

const rest = new REST({ version: "9" }).setToken(GetToken);

rest.put(Routes.applicationGuildCommands(GetClientId, GetGuildId), { body: commands })
	.then(() => console.log("Successfully registered application commands."))
	.catch(console.error);

const client = new Client({ intents: [Intents.FLAGS.GUILDS] });

// When the client is ready, run this code (only once)
client.once("ready", () => {
	console.log("Ready!");
});

// Login to Discord with your client's token
client.login(GetToken);

client.on("interactionCreate", async (interaction) => {
	if (!interaction.isCommand()) return;

	const { commandName } = interaction;

	if (commandName === "ping") {
		await interaction.reply("Pong!");
	} else if (commandName === "server") {
		await interaction.reply(%s);
	}
});`, "`Server name: ${interaction.guild.name}\nTotal members: ${interaction.guild.memberCount}`")
}

func Resources() string {
	return `# Botway Discord (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [A powerful JavaScript library for interacting with the Discord API](https://github.com/discordjs/discord.js)
- [Discord.js Website](https://discord.js.org)
- [Discord.js Docs](https://discord.js.org/#/docs)
- [Discord.js Guide](https://discordjs.guide)
- [Discord Server](https://discord.gg/djs)

big thanks to [**@discordjs**](https://github.com/discordjs) org`
}
