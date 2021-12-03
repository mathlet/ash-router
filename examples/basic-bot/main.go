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

	// Add default router handlers to session. This will add all global commands
	// and it will also add all guild commands to every guild your bot is in. If you
	// want to avoid this checkout the advanced-bot directory.
	router.StartWithDefault(s, nil)

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
