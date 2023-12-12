package events

import (
	"strings"

	"agathabot/internal/guild"
	"agathabot/internal/player"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(client *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == client.State.User.ID {
		return
	}

	g := m.Message.GuildID
	if _, ok := guild.Guilds[g]; ok {
		c := m.Message.ChannelID
		if *guild.Guilds[g].ChannelDM == c {
			//guild dm chat
			if strings.Contains(strings.ToLower(m.Message.Content), "execute order 66") {
				client.ChannelMessageSend(c, "It will be done my lord")
				client.ChannelMessageSend(c, "︻╦╤─ - - -      o")
				client.ChannelMessageSend(c, "︻╦╤─ - - -      T")
				client.ChannelMessageSend(c, "︻╦╤─ - - -      ^")
			}
			if strings.Contains(strings.ToLower(m.Message.Content), " of darth plagueis") {
				client.ChannelMessageSend(c, "Did you ever hear the tragedy of Darth Plagueis the Wise?")
				client.ChannelMessageSend(c, "I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis... was a Dark Lord of the Sith so powerful and so wise, he could use the Force to influence the midi-chlorians... to create... life. He had such a knowledge of the dark side, he could even keep the ones he cared about... from dying.")
				client.ChannelMessageSend(c, "The dark side of the Force is a pathway to many abilities... some consider to be unnatural.")
				client.ChannelMessageSend(c, "He became so powerful, the only thing he was afraid of was... losing his power. Which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew. Then his apprentice killed him in his sleep. It's ironic. He could save others from death, but not himself.")
			}
			if strings.Contains(strings.ToLower(m.Message.Content), " is evil") {
				client.ChannelMessageSend(c, "From my point of view, the Jedi are evil!")
			}
			if strings.Contains(strings.ToLower(m.Message.Content), " i can't ") {
				client.ChannelMessageSend(c, "Because of Obi-Wan?")
			}
			if strings.Contains(strings.ToLower(m.Message.Content), "leave everything") {
				client.ChannelMessageSend(c, "Don't you see? We don't have to run away anymore. I have brought peace to the Republic. I am more powerful than the Chancellor. I... I can overthrow him. And together, you and I can rule the galaxy... make things the way we want them to be!")
			}
			if strings.Contains(strings.ToLower(m.Message.Content), "high ground") {
				client.ChannelMessageSend(c, "Don't underestimate my power!")
			}
			return
		} else {
			//not dm channel
			return
		}
	} else {
		p := m.Message.Author.ID
		if ply, ok := player.Players[p]; !ok {
			//not an active player
			return
		} else {
			//player chatting in dm
			ply.HandleAction(client, m.Message.Content)
		}
	}
}
