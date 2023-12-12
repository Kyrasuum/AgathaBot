package commands

import (
	"fmt"

	"agathabot/internal/guild"
	"agathabot/internal/player"

	"github.com/bwmarrin/discordgo"
)

func status(client *discordgo.Session, i *discordgo.InteractionCreate) {
	g := guild.Guilds[i.Interaction.GuildID]

	status := fmt.Sprintf("Current Players(%d):", len(g.Players))
	for p, nick := range g.Players {
		status += fmt.Sprintf("\n\t%+v", nick)
		if ply, ok := player.Players[p]; ok {
			status += fmt.Sprintf("-> %+v", ply.Name)
		}
	}
	client.ChannelMessageSend(i.Interaction.ChannelID, status)
}
