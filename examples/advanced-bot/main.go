package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/mathlet/ashrouter"
)

var token string = "your token goes here"

func main() {
	// Create new discordgo session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	// Set session intents
	s.Identify.Intents = discordgo.IntentsAll

	// Create new router
	router := ashrouter.NewRouter()

	// Add commands to router (you can also do this manually)
	router.AddGlobalCommands(Ping)
	router.AddGuildCommands(Echo)

	// Add custom handlers to create commands
	s.AddHandler(func(ses *discordgo.Session, g *discordgo.GuildCreate) {
		// Add commands to specific guild (you don't have to do this inside a handler)
		if g.Name == "your guild name" {
			router.CreateGuildCommands(ses, g.ID)
		}
	})

	// You can also add the default command executor or you can create your own.
	// This will check InteractionCreate events for your commands.
	router.DefaultCommandExecutor(s, nil)

	// Open discordgo session
	err = s.Open()
	if err != nil {
		panic(err)
	}

	// Wait for CTRL+C to close session
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Print("Shutting down gracefully...")
	if err := s.Close(); err != nil {
		panic(err)
	}
}
