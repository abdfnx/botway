<div align="center">
  <h1>Botway-Swift</h1>
	<p>
		Swift client package for Botway
	</p>
</div>

## Usage

> after creating a new swift botway project, you need to use your tokens to connect with your bot.

```swift
import Foundation
import TelegramBotSDK

let bot = TelegramBot(token: GetToken())
let router = Router(bot: bot)

router["greet"] = { context in
    guard let from = context.message?.from else { return false }
    context.respondAsync("Hello, \(from.firstName)!")

    return true
}

router[.newChatMembers] = { context in
    guard let users = context.message?.newChatMembers else { return false }

    for user in users {
        guard user.id != bot.user.id else { continue }
        context.respondAsync("Welcome, \(user.firstName)!")
    }

    return true
}

while let update = bot.nextUpdateSync() {
	try router.process(update: update)
}

fatalError("Server stopped due to error: \(bot.lastError)")
```
