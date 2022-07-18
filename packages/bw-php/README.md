<div align="center">
  <h1>bw-php</h1>
	<p>
		PHP client library for Botway.
	</p>
	<br />
	<p></p>
</div>

```bash
composer require abdfnx/bw-php
```

## Usage

> after creating a new php botway project, you need to use your tokens to connect with your bot.

```php
<?php

use Discord\Discord;
use Discord\Parts\Channel\Message;
use Botway\Botway;

$botConfig = new Botway();

// Create a $discord BOT
$discord = new Discord([
    "token" => $botConfig->GetToken(),
]);

$discord->on("ready", function (Discord $discord) {
    // Listen for messages
    $discord->on("message", function (Message $message, Discord $discord) {
        // If message is from a bot
        if ($message->author->bot) {
            // Do nothing
            return;
        }

        // If message is "ping"
        if ($message->content == "ping") {
            // Reply with "pong"
            $message->reply("pong");
        }
    });
});

// Start the Bot (must be at the bottom)
$discord->run();
```
