package guild

import (
	"github.com/bwmarrin/discordgo"
)

type GuildInfo struct {
	ChannelDM      *string
	RegisteredCmds []*discordgo.ApplicationCommand
}

var Guilds = map[string]GuildInfo{}
