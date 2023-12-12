package commands

import (
	"agathabot/internal/guild"

	"github.com/bwmarrin/discordgo"
)

var ()

func setup(client *discordgo.Session, i *discordgo.InteractionCreate) {
	// Find the channel that the message came from.
	c, err := client.State.Channel(i.Interaction.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	// Find the guild for that channel.
	g, err := client.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}

	err = client.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "AgathaBot will now DM from this channel",
		},
	})
	if err == nil {
		*guild.Guilds[g.ID].ChannelDM = c.ID
	}
}
