# Discord Ping Bot

A simple Discord bot written in Go (Golang) that responds to ping commands.  
This project demonstrates how to build a Discord bot using the [`discordgo`](https://github.com/bwmarrin/discordgo) library and handle basic commands.

---

## Features

- Responds to `!ping` command with `Pong!`
- Easy to configure via JSON config file
- Lightweight and efficient using Go concurrency
- Suitable as a starting point for more complex Discord bots

---

## Getting Started

### Prerequisites

- Go 1.18+ installed ([Download Go](https://golang.org/dl/))
- A Discord bot token ([Create a bot on Discord Developer Portal](https://discord.com/developers/applications))
- Basic knowledge of Go and Discord bots

### Installation

1. Clone the repository:

git clone https://github.com/KDhruv3/Discord-Ping.git
cd Discord-Ping


2. Create a `config.json` file in the project root with the following structure:

{
"Token": "YOUR_DISCORD_BOT_TOKEN",
"BotPrefix": "!"
}


3. Install dependencies:

go mod tidy

### Running the Bot

Run the bot with:
go run main.go


The bot will connect to Discord and listen for commands prefixed with the configured `BotPrefix`.

---

## Usage

- In your Discord server where the bot is added, type:

!ping

- The bot will reply with:

Pong!

---

## Project Structure

.
├── main.go # Entry point of the bot
├── config.go # Configuration handling
├── handlers.go # Command and event handlers
├── go.mod # Go module file
├── go.sum # Go dependencies checksum
└── config.json # Bot configuration file (not committed)



---

## About `discordgo`

This project uses [`discordgo`](https://github.com/bwmarrin/discordgo), a popular Go package that provides bindings for the Discord API. It simplifies the process of building Discord bots by handling WebSocket connections, events, and REST API calls.

- **Repository:** [https://github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)  
- **Features:**  
  - Real-time event handling (messages, reactions, voice, etc.)  
  - Voice support (note: voice features may have some known issues, see the repo for details)  
  - Easy command handling  
  - Active community and ongoing maintenance  

### Why `discordgo`?

- It is the de facto Go library for Discord bots, widely used and well-supported.  
- Provides a clean and idiomatic Go interface to Discord’s API.  
- Enables rapid development of Discord bots with concurrency and event-driven design.

---

## Example usage snippet

import "github.com/bwmarrin/discordgo"

func main() {
dg, err := discordgo.New("Bot " + "YOUR_BOT_TOKEN")
if err != nil {
panic(err)
}
dg.AddHandler(messageCreate)

err = dg.Open()
if err != nil {
    panic(err)
}

// Keep the program running until terminated
select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
if m.Content == "!ping" {
s.ChannelMessageSend(m.ChannelID, "Pong!")
}
}

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [discordgo](https://github.com/bwmarrin/discordgo) - Go package for Discord API  
- [Discord Developer Portal](https://discord.com/developers) - For creating and managing your bot

---

If you have any questions or need help, feel free to reach out!

