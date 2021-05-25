package main

import "github.com/bwmarrin/discordgo"

func help(s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSend(m.ChannelID, "https://github.com/Tch1b0/MoGo")
}