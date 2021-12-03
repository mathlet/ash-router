package ashrouter

import "github.com/bwmarrin/discordgo"

// Context is a command context passed into all command handlers.
type Context struct {
	// Discordgo session
	*discordgo.Session
	// Interaction of command
	*discordgo.Interaction
	// An interface for any data you want to pass to your commands
	Vars interface{}
}
