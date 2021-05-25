package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	linker "github.com/Tch1b0/MoGo/linker"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("Please create a token.txt file and store the key of your bot in there.")
		return
	}
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", token))

	if err != nil {
		fmt.Println("Couldn't create Discord session: ", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Couldn't open a Discord session: ", err)
	}

	fmt.Println("MoGo is now on Air!  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "Just looking around")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "$") {
		msg := strings.Split(m.Content, " ")
		rawreg := `^https?:\/\/(([a-z0-9]){0,}\.)?([a-z0-9]){2,63}\.[a-z]{2,}(\/[\s\S]{0,}?){0,}$` // Yes, I wrote this by myself and yes, it took me a few hours
		reg, err := regexp.Compile(rawreg)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i:=0; i < len(msg); i++ {
			
			matched := reg.MatchString(msg[i]) 
			fmt.Println(matched)
			if err != nil {
				fmt.Println(err)
				return
			}
			if matched && len(msg[i]) > 35 {
				l := linker.Linker{Link: msg[i]}
				l, err = l.Create()
				if err != nil {
					fmt.Println(err)
					return
				}
				s.ChannelMessageSend(
					m.ChannelID, 
					fmt.Sprintf("Shorter Link: https://ls.johannespour.de/%s", l.Short),
				)
			}
		}
		return
	}

	m.Content = strings.Replace(m.Content, "$", "", 1)

	c := strings.Split(m.Content, " ")

	if c[0] == "short" {
		if len(c) < 2{
			s.ChannelMessageSend(m.ChannelID, "Link is missing")
			return
		}

		l := linker.Linker{Link: c[1]}
		l, err := l.Create()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("LINK: %s, SHORT: %s, TOKEN: %s\n", l.Link, l.Short, l.Token)
		s.ChannelMessageSend(
			m.ChannelID, 
			fmt.Sprintf("Your shortcut is:%s\nThe corrosponding link is: https://ls.johannespour.de/%s", l.Short, l.Short),
		)
	}
}