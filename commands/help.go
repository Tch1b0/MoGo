package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate){
	e := discordgo.MessageEmbed{}
	e.Description = fmt.Sprintf(
		"Hi! I am **%s** and do you know what I don't like? Too long links.\n"+
		"But that is what I was made for!\n"+
		"I am here to **automatically** shorten your links!\n\n"+
		"And here is how it works:\n"+
		"You can just write messages ***like you are used to***.\n"+
		"If there is a link in your message, and it's longer than my shortcut,\n **I am making it shorter for you!**",
		s.State.User.Username,
	)
	e.Color = 0xFF80ED
	s.ChannelMessageSendEmbed(
		m.ChannelID, 
		&e,
	)
}