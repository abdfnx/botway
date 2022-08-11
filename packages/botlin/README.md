<div align="center">
  <h1>Botlin</h1>
	<p>
		Kotlin client package for Botway
	</p>
</div>

## Usage

> after creating a new kotlin botway project, you need to use your tokens to connect with your bot.

```kotlin
package core

import dev.kord.core.*
import dev.kord.core.entity.*
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.gateway.*
import botway.*

suspend fun main() {
    val kord = Kord(GetToken())
...
```
