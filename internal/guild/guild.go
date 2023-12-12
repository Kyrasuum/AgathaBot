package guild

import (
	"github.com/bwmarrin/discordgo"
)

type GuildInfo struct {
	ChannelDM      *string
	Players        map[string]string
	RegisteredCmds []*discordgo.ApplicationCommand
}

var Guilds = map[string]*GuildInfo{}
