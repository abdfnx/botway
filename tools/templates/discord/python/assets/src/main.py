import discord
from botway import GetToken

intents = discord.Intents.default()

client = discord.Client(intents=intents)

@client.event
async def on_ready():
	print(f'We have logged in as {client.user}')

@client.event
async def on_message(message):
	if message.author == client.user:
		return

	if message.content.startswith('Hi'):
		await message.channel.send('Hello!')

client.run(GetToken())
