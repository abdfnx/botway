<div align="center">
  <h1>botway_dart</h1>
	<p>
		Dart client package for Botway
	</p>
	<br />
	<p>
		<img alt="Pub" src="https://img.shields.io/pub/v/botway_dart?label=pub.dev&logo=dart">
	</p>
</div>

## Usage

> after creating a new dart botway project, you need to use your tokens to connect with your bot.

```dart
import "package:nyxx/nyxx.dart";
import "package:botway_dart/botway_dart.dart";

void main() {
  var bot_config = Botway();

  final bot = NyxxFactory.createNyxxWebsocket(bot_config.Get_Token(), GatewayIntents.allUnprivileged)
    ..registerPlugin(Logging()) // Default logging plugin
    ..registerPlugin(CliIntegration()) // Cli integration for nyxx allows stopping application via SIGTERM and SIGKILl
    ..registerPlugin(IgnoreExceptions()) // Plugin that handles uncaught exceptions that may occur
    ..connect();

...
```
