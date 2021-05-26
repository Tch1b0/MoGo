package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Tch1b0/MoGo/commands"
	"github.com/bwmarrin/discordgo"
)

func ScanForLinks(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := strings.Split(m.Content, " ")
	rawreg := `^https?:\/\/(([a-z0-9]){0,}\.)?([a-z0-9]){2,63}\.[a-z]{2,}(\/[\s\S]{0,}?){0,}$` // Yes, I wrote this by myself and yes, it took me a few hours
	reg, err := regexp.Compile(rawreg)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(msg); i++ {

		matched := reg.MatchString(msg[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if matched && len(msg[i]) > 35 {
			commands.Short(s, m, msg[i], false)
		}
	}
}