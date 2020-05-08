package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_AUTHENTICATION_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		if mc.Author.ID == s.State.User.ID || !strings.HasPrefix(mc.Content, "!testbot") {
			return
		}
		text := strings.TrimLeft(mc.Content, "!testbot")
		_, err = s.ChannelMessageSend(mc.ChannelID, text)
		if err != nil {
			log.Print(err)
		}
	})
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = discord.Close()
	if err != nil {
		log.Fatal(err)
	}
}
