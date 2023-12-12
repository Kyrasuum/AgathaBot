package commands

import (
	"github.com/bwmarrin/discordgo"
)

// All commands registered
var (
	Cmds = []*discordgo.ApplicationCommand{
		{
			Name:        "setup",
			Description: "setup",
		},
	}

	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"setup": setup,
	}
)
