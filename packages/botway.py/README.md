<div align="center">
  <h1>botway.py</h1>
	<p>
		Python client package for Botway.
	</p>
	<br />
	<p>
		<img alt="PyPI" src="https://img.shields.io/pypi/v/botway.py?logo=python&style=flat-square">
	</p>
</div>

```bash
# pip
# Linux/macOS
pip3 install botway.py

# Windows
pip install botway.py

# pipenv
pipenv install botway.py
```

## Usage

> after creating a new python botway project, you need to use your tokens to connect with your bot.

```python
...
from botway import GetToken

from telegram import Update, ForceReply
from telegram.ext import Updater, CommandHandler, MessageHandler, Filters
...
def main() -> None:
	"""Start the bot."""
	# Create the Updater and pass it your bot's token.
	updater = Updater(GetToken())

	# Get the dispatcher to register handlers
	dispatcher = updater.dispatcher
...
```
