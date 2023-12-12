package commands

import (
	"agathabot/internal/player"
	"agathabot/pkg/io"

	"github.com/bwmarrin/discordgo"
)

// All commands registered
var (
	ServCmds = []*discordgo.ApplicationCommand{
		{
			Name:        "setup",
			Description: "Sets AgathaBot to DM out of the current channel",
		},
		{
			Name:        "status",
			Description: "Retrieves current status about the running game",
		},
		{
			Name:        "join",
			Description: "Lets AgathaBot know that you would like to join into the current running game",
		},
	}
	DmCmds = []*discordgo.ApplicationCommand{}

	ServHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"setup":  setup,
		"status": status,
		"join":   join,
	}
	DmHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"status": plyWrapper(nil, "Status"),
	}
)

func plyWrapper(pkg interface{}, fn string) func(client *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(client *discordgo.Session, i *discordgo.InteractionCreate) {
		p := i.Interaction.Member.User.ID
		ply := player.Players[p]

		if pkg == nil {
			pkg = ply
		}
		modmap, err := io.StructToMap(pkg)
		if err != nil {
			return
		}
		if modfunc, ok := modmap[fn]; !ok {
			return
		} else {
			switch typefunc := modfunc.(type) {
			case func(*discordgo.Session, *discordgo.InteractionCreate):
				typefunc(client, i)
			case func(*discordgo.Session, *discordgo.InteractionCreate, *player.Player):
				typefunc(client, i, ply)
			}
		}

	}
}
