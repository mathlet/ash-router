package ashrouter

import "github.com/bwmarrin/discordgo"

// A Command has an underlying discordgo.ApplicationCommand pointer, middlewares, and a handler.
type Command struct {
	// Discordgo slash command
	*discordgo.ApplicationCommand
	// The command's middlewares
	Middlewares []Middleware
	// The command's handler function
	Handler CommandHandler
}

// ApplyMiddlewares applies all of the commands middlewares to it's handler
func (c *Command) ApplyMiddlewares() CommandHandler {
	h := c.Handler
	for _, m := range c.Middlewares {
		h = m(h)
	}
	return h
}
