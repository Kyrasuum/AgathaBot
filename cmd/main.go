package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"agathabot/internal/events"
	"agathabot/internal/guild"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	client, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Printf("Client creation Error: %v", err)
		return
	}

	for _, handler := range []interface{}{
		events.ChannelDelete,
		events.GuildCreate,
		events.GuildDelete,
		events.InteractionCreate,
		events.MessageCreate,
		events.Ready,
	} {
		client.AddHandler(handler)
	}

	// Open a websocket connection to Discord and begin listening.
	err = client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer client.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	fmt.Println("Shutting down...")

	//cleanup commands from guilds
	for id, g := range guild.Guilds {
		for _, cmd := range g.RegisteredCmds {
			err := client.ApplicationCommandDelete(client.State.User.ID, id, cmd.ID)
			if err != nil {
				fmt.Printf("Cannot delete '%v' command: %v", cmd.Name, err)
			}
		}
	}
}
