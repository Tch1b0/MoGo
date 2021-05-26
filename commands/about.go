package commands

import "github.com/bwmarrin/discordgo"

func About(s *discordgo.Session, m *discordgo.MessageCreate) {
	e := discordgo.MessageEmbed{}
	e.Title = "About"
	e.Description = ""
	s.ChannelMessageSendEmbed(
		m.ChannelID,
		&e,
	)
}