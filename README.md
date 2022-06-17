<p align="center">
  <a href="https://botway.web.app" target="_blank">
    <img src="https://cdn-botway.up.railway.app/botway.svg" alt="Botway" width="300">
  </a>
</p>

> **🤖 Generate, build, handle and deploy your own bot with your favorite language, for Discord, or Telegram, or Slack.**

With botway, you can focus on your bot's logic and don't worry about the infrastructure. and we will take care of the rest.

Botway uses [Railway][rw] to host your bot code and database.

## Requirements

- [**Railway Account**][rw]

## Installation ⬇

### NPM

```bash
# npm
npm i -g botway

# yarn
yarn global add botway

# pnpm
pnpm add -g botway
```

### Using script

- Shell

```bash
curl -sL https://bit.ly/botway | bash
```

- PowerShell

```powershell
iwr -useb https://bit.ly/bw-win | iex
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

* Initialize `~/.botway`

  ```bash
  botway init
  ```

* Authenticate with [**Railway**][rw]

  ```bash
  botway login
  ```

* Open Botway TUI

  ```bash
  botway
  ```

* Create a new botway project

  ```bash
  botway new <project-name>
  ```

* Manage your bot tokens

  ```bash
  botway tokens <command> [flags] <project-name>
  ```

* Start running your bot

  ```bash
  # Under the project directory
  botway start
  ```

* Manage your bot database

  ```bash
  # Under the project directory
  botway database <command>
  ```

* Deploy and upload project to [**Railway**][rw] from the current directory

  ```bash
  # Under the project directory
  botway deploy
  ```

* Run a local command using variables from the active environment

  ```bash
  # Under the project directory
  botway run <command>
  ```

## Roadmap

> You can see the [**Roadmap**](https://github.com/users/abdfnx/projects/10)

## Keyboard shortcuts

- <kbd>Up</kbd>: **Move up**
- <kbd>Down</kbd>: **Move down**
- <kbd>Tab</kbd>: **Switch windows**
- <kbd>Ctrl+O</kbd>: **Open bot project at Railway**
- <kbd>Esc</kbd>: **Reset**
- <kbd>Ctrl+Q</kbd>: **Quit**

### Technologies Used in Botway

- [**Railway API**][rw]
- [**Charm**](https://charm.sh)
- [**Cobra**](https://github.com/spf13/cobra)
- [**Viper**](https://github.com/spf13/viper)
- [**GJson**](https://github.com/tidwall/gjson)
- [**Termenv**](https://github.com/muesli/termenv)
- [**Boa**](https://github.com/elewis787/boa)

## Special thanks ❤

Thanks to [**@charmbracelet**](https://github.com/charmbracelet) for thier awesome TUI libraries 🏗.

Also thanks to [**@railwayapp**](https://github.com/railwayapp) for amazing cloud services ☁️.

### License

botway is licensed under the terms of [MIT](https://github.com/abdfnx/botway/blob/main/LICENSE) license.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=abdfnx/botway&type=Date)](https://star-history.com/#abdfnx/botway)

[rw]: https://railway.app
