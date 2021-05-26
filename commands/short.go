package commands

import (
	"fmt"

	"github.com/Tch1b0/MoGo/utils"
	"github.com/bwmarrin/discordgo"
)


func Short(s *discordgo.Session, m *discordgo.MessageCreate, link string) {
	l := utils.Linker{Link: link}
	l, err := l.Create()
	if err != nil {
		fmt.Println(err)
	}
	s.ChannelMessageSend(
		m.ChannelID, 
		fmt.Sprintf("Your shortcut is:%s\nThe corrosponding link is: https://ls.johannespour.de/%s", l.Short, l.Short),
	)
}