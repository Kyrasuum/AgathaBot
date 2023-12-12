package events

import (
	"fmt"

	"agathabot/internal/commands"
	"agathabot/internal/guild"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(client *discordgo.Session, i *discordgo.InteractionCreate) {
	//check if in registered guild
	if _, ok := guild.Guilds[i.Interaction.GuildID]; ok {
		//respond to interaction
		if h, ok := commands.ServHandlers[i.ApplicationCommandData().Name]; ok {
			h(client, i)
		} else {
			fmt.Printf("invalid command\n")
			fmt.Printf("%+v\n", i)
		}
	} else {
		//respond to interaction on dm
		if h, ok := commands.DmHandlers[i.ApplicationCommandData().Name]; ok {
			h(client, i)
		} else {
			fmt.Printf("invalid command\n")
			fmt.Printf("%+v\n", i)
		}
	}
}
