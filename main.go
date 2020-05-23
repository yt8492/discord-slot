package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/ini.v1"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Config struct {
	Prefix string
	DiscordAuthenticationToken string
}

var Conf Config

func init() {
	c, _ := ini.Load("config.ini")
	Conf = Config{
		Prefix: c.Section("config").Key("prefix").MustString("$"),
		DiscordAuthenticationToken: c.Section("config").Key("token").String(),
	}
}

func main() {
	discord, err := discordgo.New("Bot " + Conf.DiscordAuthenticationToken)
	if err != nil {
		log.Fatal(err)
	}
	prefix := Conf.Prefix
	discord.AddHandler(func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		log.Println(mc.Content)
		if mc.Author.ID == s.State.User.ID || !strings.HasPrefix(mc.Content, prefix) {
			return
		}
		command := strings.TrimPrefix(mc.Content, prefix)
		var g *discordgo.Guild
		g, err = s.Guild(mc.GuildID)
		if err != nil {
			log.Println(err)
			return
		}
		switch command {
		case "slot":
			emojis := getGuildEmojis(g)
			slot := make([]string, 3)
			for i := 0; i < 3; i++ {
				slot[i] = emojis[rand.Intn(len(emojis))]
			}
			var result string
			if len(slot) > 0 {
				result = strings.Join(slot, "")
			} else {
				result = "This guild has no emoji."
			}
			_, err = s.ChannelMessageSend(mc.ChannelID, result)
		case "list":
			emojis := getGuildEmojis(g)
			_, err = s.ChannelMessageSend(mc.ChannelID, strings.Join(emojis, "\n"))
		}
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

func getGuildEmojis(g *discordgo.Guild) []string {
	emojis := make([]string, len(g.Emojis))
	for i, emoji := range g.Emojis {
		emojis[i] = fmt.Sprintf("<:%s:%s>", emoji.Name, emoji.ID)
	}
	return emojis
}
