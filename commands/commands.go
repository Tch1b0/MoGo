package commands

import (
	"github.com/bwmarrin/discordgo"
)

func CommandList(s *discordgo.Session, m *discordgo.MessageCreate) {
	e := discordgo.MessageEmbed{}
	e.Title = "Commands"
	e.Color = 0x0BEEF0
	e.Fields = append(
		e.Fields,
		&discordgo.MessageEmbedField{
			Name: "$short",
			Value: "Shorten a certain link",
		},
		&discordgo.MessageEmbedField{
			Name: "$help",
			Value: "Get a short summary what the bot is used for",
		},
		&discordgo.MessageEmbedField{
			Name: "$about",
			Value: "Get further information about how the bot was made",
		},
		&discordgo.MessageEmbedField{
			Name: "$commands",
			Value: "Get a list with all commands available",
		},
	)
	s.ChannelMessageSendEmbed(
		m.ChannelID,
		&e,
	)
}