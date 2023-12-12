package events

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(client *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == client.State.User.ID {
		return
	}
}
