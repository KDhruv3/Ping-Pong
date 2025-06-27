# Discord Ping Bot

A simple Discord bot written in Go that responds to the `!ping` command with `Pong!`.

## Features

- Responds to `!ping` command
- Built using the `discordgo` Go library
- Easy to configure with a JSON config file

## Getting Started

### Prerequisites

- Go installed (version 1.18 or higher)
- A Discord bot token from the Discord Developer Portal

### Installation

1. Clone the repository:

    git clone https://github.com/KDhruv3/Ping-Pong.git
    
    cd Ping-Pong


2. Create a `config.json` file with your bot token and prefix:

    {
        "Token": "YOUR_DISCORD_BOT_TOKEN",
        "BotPrefix": "!"
    }

3. Install dependencies:
    go mod tidy


### Running the Bot

Run the bot with:

    go run main.go


## Usage

Type `!ping` in your Discord server where the bot is added, and it will reply with `Pong!`.

