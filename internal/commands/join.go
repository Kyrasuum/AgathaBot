package commands

import (
	"fmt"

	"agathabot/internal/character"
	"agathabot/internal/guild"
	"agathabot/internal/player"

	"github.com/bwmarrin/discordgo"
)

func join(client *discordgo.Session, i *discordgo.InteractionCreate) {
	p := i.Interaction.Member.User.ID
	//We create the private channel with the user who sent the message.
	channel, err := client.UserChannelCreate(p)
	if err != nil {
		fmt.Printf("error creating channel: %+v\n", err)
		client.ChannelMessageSend(i.ChannelID, "Something went wrong while sending the DM!")
	}

	_, err = client.ChannelMessageSend(channel.ID, "Welcome to Age of Agatha!")
	if err != nil {
		fmt.Printf("error sending message: %+v\n", err)
		client.ChannelMessageSend(i.Interaction.ChannelID, "Failed to send you a DM. Did you disable DM in your privacy settings?")
	}

	//initialize player data
	g := guild.Guilds[i.Interaction.GuildID]
	ply := player.Players[p]
	if ply == nil {
		ply = &player.Player{
			Name:      i.Interaction.Member.Nick,
			ChannelDM: channel.ID,

			Avatar: &character.Character{},
		}
		player.Players[p] = ply
	}
	g.Players[p] = i.Interaction.Member.Nick

	//create character if nil
	if ply.Avatar == nil {
		ply.Avatar = &character.Character{}
	}

	//respond in guild channel
	client.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Welcome %+v to Age of Agatha!", i.Interaction.Member.Nick),
		},
	})
}
