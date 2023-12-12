package player

import (
	"agathabot/internal/character"

	"github.com/bwmarrin/discordgo"
)

type Player struct {
	Name       string
	ChannelDM  string
	CurrAction *Action

	Avatar *character.Character
}

type Event interface{}
type Action struct {
	Prompt       string
	Options      map[string]Event
	Descriptions map[string]string
}

var Players map[string]*Player

func (ply *Player) HandleAction(client *discordgo.Session, msg string) {
	if ply.CurrAction != nil {
		if NextAction, ok := ply.CurrAction.Options[msg]; ok {
			//valid action
			switch act := NextAction.(type) {
			//sub action
			case *Action:
				ply.DisplayAction(client, act)
				ply.CurrAction = act
			//event
			default:
				ply.HandleEvent(client, act)
			}
		} else {
			//invalid action
			client.ChannelMessageSend(ply.ChannelDM, "Action not recognized")
		}
	}
}

func (ply *Player) DisplayAction(client *discordgo.Session, a *Action) {
	client.ChannelMessageSend(ply.ChannelDM, a.Prompt)
	client.ChannelMessageSend(ply.ChannelDM, " ")
	client.ChannelMessageSend(ply.ChannelDM, "Please select from the following options:")
	for option, desc := range a.Descriptions {
		client.ChannelMessageSend(ply.ChannelDM, option+": "+desc)
	}
}

func (ply *Player) HandleEvent(client *discordgo.Session, ev Event) {}
