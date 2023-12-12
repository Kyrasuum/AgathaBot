package events

import (
	"fmt"

	"agathabot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(client *discordgo.Session, i *discordgo.InteractionCreate) {
	//respond to interaction
	if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
		h(client, i)
	} else {
		fmt.Printf("invalid command\n")
		fmt.Printf("%+v\n", i)
	}
}
