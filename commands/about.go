package commands

import (
	"github.com/bwmarrin/discordgo"
)

func About(s *discordgo.Session, m *discordgo.MessageCreate) {
	e := discordgo.MessageEmbed{}
	e.Title = "About"
	e.Description = (
		"My developer is **Tch1b0**\n and you can find my **sourcecode** on GitHub:\n"+
		"https://github.com/Tch1b0/MoGo")			
	s.ChannelMessageSendEmbed(
		m.ChannelID,
		&e,
	)
}