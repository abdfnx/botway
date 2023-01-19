<p align="center">
  <a href="https://botway.deno.dev" target="_blank">
    <img src="https://cdn-botway.deno.dev/botway.svg" alt="Botway" width="300">
  </a>
</p>

> **🤖 Generate, build, handle and deploy your own bot with your favorite language, for Discord, or Telegram, or Slack, or even Twitch.**

With botway, you can focus on your bot's logic and don't worry about the infrastructure. and we will take care of the rest.

Botway uses [Railway][rw] and [Render][rnd] to host your bot code and database.

<p align="center">
  <img src="https://cdn-botway.deno.dev/screenshots/deploy.svg" alt="Botway Screenshot" width="1400">
</p>

## Requirements

- [**Railway Account**][rw]
- [**Render Account**][rnd]

## Installation ⬇

### Using script

- Shell

```bash
curl -sL https://dub.sh/botway | bash
```

- PowerShell

```powershell
irm https://dub.sh/bw-win | iex
```

**then restart your powershell**

### Homebrew

```
brew install abdfnx/tap/botway
```

### Scoop

```
scoop bucket add botway https://github.com/abdfnx/botway
scoop install botway
```

## Usage

- Initialize `~/.botway`

  ```bash
  botway init
  ```

- Authenticate with your favorite host service

  ```bash
  # railway
  botway login railway

  # render
  botway login render
  ```

- Open Botway TUI

  ```bash
  botway
  ```

- Create a new botway project

  ```bash
  botway new <project-name>
  ```

- Manage your bot tokens

  ```bash
  botway tokens <command> [flags] <project-name>
  ```

- Start running your bot

  ```bash
  # Under the project directory
  botway start
  ```

- Manage your bot database

  ```bash
  # Under the project directory
  botway database <command>
  ```

- Deploy and upload project from the current directory

  ```bash
  # Under the project directory
  botway deploy
  ```

- Execute a local command using variables from the active environment

  ```bash
  # Under the project directory
  botway exec <command>
  ```

## Roadmap

> You can see the [**Roadmap**](https://github.com/users/abdfnx/projects/10)

## Keyboard shortcuts

- <kbd>Up</kbd>: **Move up**
- <kbd>Down</kbd>: **Move down**
- <kbd>Tab</kbd>: **Switch windows**
- <kbd>Ctrl+O</kbd>: **Open bot project at Host Service**
- <kbd>Esc</kbd>: **Reset**
- <kbd>Ctrl+Q</kbd>: **Quit**

### Technologies Used in Botway

- [**Railway API**][rw]
- [**Render Rest API**][rnd]
- [**Charm**](https://charm.sh)
- [**Cobra**](https://github.com/spf13/cobra)
- [**Viper**](https://github.com/spf13/viper)
- [**GJson**](https://github.com/tidwall/gjson)
- [**Termenv**](https://github.com/muesli/termenv)
- [**Boa**](https://github.com/elewis787/boa)

## Special thanks ❤

Thanks to [**@charmbracelet**](https://github.com/charmbracelet) for their awesome TUI libraries 🏗.

Also thanks to [**@railwayapp**](https://github.com/railwayapp) and [**@renderinc**](https://github.com/renderinc) for their amazing cloud and host services ☁️.

### License

botway is licensed under the terms of [MIT](https://github.com/abdfnx/botway/blob/main/LICENSE) license.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=abdfnx/botway&type=Date)](https://star-history.com/#abdfnx/botway)

[rw]: https://railway.app
[rnd]: https://render.com
