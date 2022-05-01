package python

import "github.com/abdfnx/botway/tools/templates"

func MainPyContent() string {
	return templates.Content("slack", "python", "src/main.py", "")
}

func Flake8Content() string {
	return templates.Content("slack", "python", ".flake8", "")
}

func ProcfileContent() string {
	return templates.Content("slack", "python", "Procfile", "")
}

func Resources() string {
	return `# Botway Slack (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/6)

## API

- [A lightweight framework to quickly develop modern Slack bots](https://github.com/Ambro17/slackify)
- [Slackify Website](https://ambro17.github.io/slackify)

## Examples

- [A collection of examples written with Slackify](https://github.com/Ambro17/slackify/tree/master/examples)

big thanks to [**@Ambro17**](https://github.com/Ambro17)`
}
