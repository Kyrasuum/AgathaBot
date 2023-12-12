package events

import (
	"fmt"

	"agathabot/internal/commands"
	"agathabot/internal/guild"
	"agathabot/internal/player"

	"github.com/bwmarrin/discordgo"
)

func GuildCreate(client *discordgo.Session, g *discordgo.GuildCreate) {
	//register guild
	channeldm := ""
	guild.Guilds[g.Guild.ID] = &guild.GuildInfo{
		ChannelDM:      &channeldm,
		Players:        map[string]*player.Player{},
		RegisteredCmds: make([]*discordgo.ApplicationCommand, len(commands.ServCmds)),
	}

	//register commands for guild
	for i, cmd := range commands.ServCmds {
		reg, err := client.ApplicationCommandCreate(client.State.User.ID, g.Guild.ID, cmd)
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v", cmd.Name, err)
			guild.Guilds[g.Guild.ID].RegisteredCmds[i] = nil
		} else {
			guild.Guilds[g.Guild.ID].RegisteredCmds[i] = reg
		}
	}
}

func GuildDelete(client *discordgo.Session, g *discordgo.GuildDelete) {
	//remove registered commands
	for _, cmd := range guild.Guilds[g.Guild.ID].RegisteredCmds {
		err := client.ApplicationCommandDelete(client.State.User.ID, g.Guild.ID, cmd.ID)
		if err != nil {
			fmt.Printf("Cannot delete '%v' command: %v", cmd.Name, err)
		}
	}
}
