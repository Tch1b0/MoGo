package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate){
	answer := fmt.Sprintf(
		"Hi! I am %s and you know what I don't like? Too long links.\n"+
		"But that is what I was made for!\n"+
		"I am here to **automatically** shorten your links!\n\n"+
		"Here you can get furhter information how I was made:\n"+
		"https://github.com/Tch1b0/MoGo", 
		s.State.User.Username,
	)
	s.ChannelMessageSend(
		m.ChannelID, 
		answer,
	)
}