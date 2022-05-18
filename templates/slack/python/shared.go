package python

import "github.com/abdfnx/botway/templates"

func MainPyContent() string {
	return templates.Content("slack/python/assets/src/main.py", "")
}

func Resources() string {
	return `# Botway Slack (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/31)

## API

- [A framework to build Slack apps using Python](https://github.com/slackapi/bolt-python)
- [Website](https://slack.dev/bolt-python)

## Examples

- [A collection of examples written with Slackify](https://github.com/slackapi/bolt-python/tree/main/examples)

big thanks to [**@slackapi**](https://github.com/slackapi) org`
}
