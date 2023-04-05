<h1 align="center"> Gabbro CLI 👾</h1>
<div align="center">

<p>

  [![Go Doc](https://godoc.org/github.com/akecel/gabbro?status.svg)](https://godoc.org/github.com/akecel/gabbro)
  [![Build](https://github.com/akecel/gabbro/actions/workflows/go.yml/badge.svg?branch=master&event=push)](https://github.com/Akecel/gabbro/actions/workflows/go.yml)
  [![Go Report Card](https://goreportcard.com/badge/github.com/akecel/gabbro)](https://goreportcard.com/report/github.com/akecel/gabbro)
  [![Coverage Status](https://coveralls.io/repos/github/Akecel/gabbro/badge.svg?branch=master)](https://coveralls.io/github/Akecel/gabbro?branch=master)
  ![Version](https://img.shields.io/github/v/release/Akecel/gabbro.svg)
  ![Licence](https://img.shields.io/badge/License-MIT-blue.svg)

</p>

</div>

## About Gabbro

Gabbro is a Cross-Platefrom CLI written in Go allowing a simplified interaction with the **IGDB API**.

### Supported plateforms

- Linux (`linux/amd64`, `linux/arm64`)
- MacOS (`darwin/amd64`, `darwin/arm64`)
- Windows (`windows/amd64`, `windows/arm64`)

### Features

- Get informations about a game by searching by it name :
  - Name, Description, Companies, Genres, Themes, DLCs, Rating, etc.
- Get informations about a character by searching by it name :
  - Name, Description, Games, Genre, Species, etc.
- Get information about the different companies involved in a game by searching by it name :
  - Name, Description, etc.

## Installation

### Download

To get started, you must download the lastest published binary of the application, make sure to download the file corresponding to your platform :

- Linux amd64 : `gabbro-linux-amd64`
- Linux arm64 : `gabbro-linux-arm64`
- MacOS amd64 : `gabbro-darwin-amd64`
- MacOS arm64 : `gabbro-darwin-arm64`
- Windows amd64 : `gabbro-windows-amd64.exe`
- Windows arm64 : `gabbro-windows-arm64.exe`

**[Download Gabbro](https://github.com/akecel/gabbro/releases/latest/)**

Note that for MacOS, you may have to download the binay with `curl` to use it :
```bash
# amd64
curl https://github.com/Akecel/gabbro/releases/download/latest/gabbro-darwin-amd64

# arm64
curl https://github.com/Akecel/gabbro/releases/download/latest/gabbro-darwin-arm64
```

Then, rename this file `gabbro` (or `gabbro.exe` for Windows)

> If you want to install Gabbro as a Go library, you can use `go install github.com/akecel/gabbro@latest`, note that the root of the configuration file will be your `$GOBIN`. You can also skip the next step and go directly to [Configuration](#Configuration).

### Unix System (MacOS / Linux)

Create a `bin` directory in your `$HOME` if it does not already exist:

```bash
mkdir $HOME/bin
```

Put the binary in your `$HOME/bin` directory and add `$HOME/bin` to your `$PATH`:

```bash
export PATH="$HOME/bin:$PATH"
```

Finally, give the right permissions to the binary:

```bash
chmod +x bin/gabbro
```

### Windows NT

TODO

### Configuration

You will now have to create crendentials for the **IGDB API** that you will put in the configuration file.

To do this, follow these steps :
- Sign Up with [Twitch](https://dev.twitch.tv/login) for a free account.
- Ensure you have Two Factor Authentication [enabled](https://www.twitch.tv/settings/security).
- Register your application in the [Twitch Developer Portal](https://dev.twitch.tv/console/apps/create).
- [Manage](https://dev.twitch.tv/console/apps) your newly created application.
- Generate a Client Secret by pressing [New Secret].
- Take note of the Client ID and Client Secret.

Once this is done, you can create a `gabbro.yaml` file in the same directory where you placed the `gabbro` binary and fill in the necessary credentials to use the IGDB API inside:

```yaml
client-id: "your-client-id"
access-token: "your-access-token"
```

You can verify that the CLI is properly installed and configured by using this command:

```bash
❯ gabbro

Usage:
  gabbro [command]

Available Commands:
  game        Get game informations
  character   Get character informations
  companies   Get information about the companies involved in a game
  help        Help about any command

Flags:
  -h, --help   help for gabbro
  -i, --image  print covers, logos and mugshots in terminal

Use "gabbro [command] --help" for more information about a command.
```

You are now ready to go ! 🚀

## Usage

### Game command

Use the `game` command to get informations about a specific game :
```bash
❯ gabbro game [Game Name] [flags]
```

```bash
Flags:
  -h, --help    help for game
  -i, --image   print covers in terminal
```

### Character command

Use the `character` command to get informations about a specific character :
```bash
❯ gabbro character [Character Name] [flags]
```

```bash
Flags:
  -h, --help    help for character
  -i, --image   print mugshot in terminal
```

### Companies command

Use the `companies` command to get informations about companies involved in a game :
```bash
❯ gabbro companies [Game Name] [flags]
```

```bash
Flags:
  -h, --help    help for companies
  -i, --image   print logos in terminal
  -o, --only string   print only one company based on the role [developer, publisher, porting, support]
```

> Note that for the --only flag, you can chain the values to get several results, for example: developer,support

### Help command

Display all available commands :
```bash
❯ gabbro help
```

Get help of a specific command :
```bash
❯ gabbro [command] --help
```

## Troubleshooting

- Images, being directly imported from IGDB can be badly displayed as for example being black on black background and thus being invisible in the terminal (especially for logos).
- Commands requiring several chained call api (`game`, `companies`) may take a few seconds to run.

## Contributing

Please read [CONTRIBUTING.md](https://github.com/Akecel/gabbro/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* [**Akecel**](https://github.com/Akecel) - *Project author*

See also the list of [contributors](https://github.com/Akecel/gabbro/graphs/contributors) who participated in this project.

## Licence

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the [LICENSE.md](https://github.com/Akecel/gabbro/blob/master/LICENSE) file for details.

## Show your support

Give a ⭐️ if this project helped you!
