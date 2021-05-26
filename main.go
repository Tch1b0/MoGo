package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Tch1b0/MoGo/commands"
	utils "github.com/Tch1b0/MoGo/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	tokenraw, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("Please create a token.txt file and store the key of your bot in there.")
		return
	}
	token := strings.Replace(string(tokenraw), "\n", "", 1)
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

func ready(s *discordgo.Session, event *discordgo.Ready) { // Called when the bot is ready
	s.UpdateGameStatus(0, "$commands")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) { // Called when a message is received

	if m.Author.ID == s.State.User.ID {
		return
	}

	c := strings.Split(m.Content, " ")  // Split the message so the arguments can be properly be read

	switch c[0] {
		case "$short":
			if len(c) < 2{
				s.ChannelMessageSend(m.ChannelID, "Link is missing")
				return
			}
			commands.Short(s, m, c[1])

		case "$help":
			commands.Help(s, m)

		case "$commands":
			commands.CommandList(s, m)

		default:
			utils.ScanForLinks(s, m)
	}
}