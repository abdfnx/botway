<div align="center">
  <h1>BW-PHP<h1>
	<p>
		PHP client package for Botway
	</p>
</div>

## Usage

> after creating a new php botway project, you need to use your tokens to connect with your bot.

```php
<?php

namespace MyBot;

include __DIR__ . "/../vendor/autoload.php";
include "botway.php";

use Discord\Discord;
use Discord\Parts\Channel\Message;
use Psr\Http\Message\ResponseInterface;
use React\Http\Browser;

$botConfig = new Botway();

// Create a $discord BOT
$discord = new Discord([
    "token" => $botConfig->GetToken(),
]);
...
```
