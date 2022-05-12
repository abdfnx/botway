const { SlashCommandBuilder } = require("@discordjs/builders");
const { REST } = require("@discordjs/rest");
const { Routes } = require("discord-api-types/v9");
const { GetToken, GetAppId, GetGuildId } = require("botway.js");
const { Client, Intents } = require("discord.js");

const commands = [
  new SlashCommandBuilder()
    .setName("ping")
    .setDescription("Replies with pong!"),
  new SlashCommandBuilder()
    .setName("server")
    .setDescription("Replies with server info!"),
].map((command) => command.toJSON());

const rest = new REST({ version: "9" }).setToken(GetToken());

rest
  // how I can get the id of my server? the answer in this github discussion https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737
  .put(
    Routes.applicationGuildCommands(GetAppId(), GetGuildId("YOUR_SERVER_NAME")),
    {
      body: commands,
    }
  )
  .then(() => console.log("Successfully registered application commands."))
  .catch(console.error);

const client = new Client({ intents: [Intents.FLAGS.GUILDS] });

// When the client is ready, run this code (only once)
client.once("ready", () => {
  console.log("Ready!");
});

// Login to Discord with your client's token
client.login(GetToken());

client.on("interactionCreate", async (interaction) => {
  if (!interaction.isCommand()) return;

  const { commandName } = interaction;

  if (commandName === "ping") {
    await interaction.reply("Pong!");
  } else if (commandName === "server") {
    await interaction.reply(
      `Server name: ${interaction.guild.name}\nTotal members: ${interaction.guild.memberCount}`
    );
  }
});
