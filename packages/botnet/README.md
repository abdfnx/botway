<div align="center">
  <h1>BotNet</h1>
	<p>
		C# client package for Botway
	</p>
	<br />
	<p>
		<img alt="Nuget" src="https://img.shields.io/nuget/v/BotNet?logo=csharp&style=flat-square">
	</p>
</div>

## Usage

> after creating a new c# botway project, you need to use your tokens to connect with your bot.

```cs
using System;
using System.Threading;
using System.Threading.Tasks;
using Discord;
using Discord.WebSocket;
using BotNet;

namespace BasicBot {
    class Program {
        private readonly DiscordSocketClient _client;

        static void Main(string[] args)
            => new Program()
                .MainAsync()
                .GetAwaiter()
                .GetResult();

        public Program() {
            _client = new DiscordSocketClient();

            _client.Log += LogAsync;
            _client.Ready += ReadyAsync;
            _client.MessageReceived += MessageReceivedAsync;
            _client.InteractionCreated += InteractionCreatedAsync;
        }

        public async Task MainAsync() {
			Core botConfig = new Core();

            await _client.LoginAsync(TokenType.Bot, botConfig.GetToken());

            await _client.StartAsync();

            await Task.Delay(Timeout.Infinite);
        }
...
```
