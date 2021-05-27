package commands

import (
	"fmt"

	"github.com/Tch1b0/MoGo/linker"
	"github.com/bwmarrin/discordgo"
)


func Short(s *discordgo.Session, m *discordgo.MessageCreate, link string, prompted bool) {
	l := linker.Linker{Link: link}
	l, err := l.Create()
	if err != nil {
		fmt.Println(err)
	}

	e := discordgo.MessageEmbed{}
	e.Title = "Link shortened"
	e.Description = fmt.Sprintf(
		"%s\n\n:arrow_down:\n\nhttps://ls.johannespour.de/%s",
		l.Link,
		l.Short,
	)
	e.Color = 0x32a8c3

	s.ChannelMessageSendEmbed(
		m.ChannelID, 
		&e,
	)

	if !l.Original() || !prompted{
		return
	}

	p := discordgo.MessageEmbed{}
	p.Title = "Shortcut"
	p.Description = fmt.Sprintf(
		"**Short:**\nhttps://johannespour.de/%s\n\n"+
		"**Token:**\n%s",
		l.Short,
		l.Token,
	)
	p.Color = 0xffff0b
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.ChannelMessageSendEmbed(
		channel.ID, 
		&p,
	)

}