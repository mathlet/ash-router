package ashrouter

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// A router stores your commands and handles InteractionCreate events
type Router struct {
	// The commands you want to add globally
	GlobalCommands []Command
	// The commands you want to add to guilds
	GuildCommands []Command
}

// NewRouter returns a pointer to a new router instance
func NewRouter() *Router {
	return &Router{}
}

// AddGlobalCommands appends your commands to router.GlobalCommands.
// commands: The commands you want to add.
func (r *Router) AddGlobalCommands(commands ...Command) {
	r.GlobalCommands = append(r.GlobalCommands, commands...)
}

// AddGuildCommands appends your commands to router.GuildCommands.
// commands: The commands you want to add.
func (r *Router) AddGuildCommands(commands ...Command) {
	r.GuildCommands = append(r.GuildCommands, commands...)
}

// CreateGuildCommands deletes old guild commands and adds all new ones. If you
// don't want to delete you can add commands manually. s: the session you want to
// create the commands on. guildID: the snowflake of the guild you want to add the
// commands to.
func (r *Router) CreateGuildCommands(s *discordgo.Session, guildID string) {
	oldCmds, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Printf("Error retrieving commands for Guild '%s': %s", guildID, err)
	}

	if len(oldCmds) > 0 {
		for _, c := range oldCmds {
			err := s.ApplicationCommandDelete(s.State.User.ID, guildID, c.ID)
			if err != nil {
				log.Printf("Error deleting old command '%s' for Guild '%s': %s", c.ID, guildID, err)
			}
		}
	}

	for _, c := range r.GuildCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, c.ApplicationCommand)
		if err != nil {
			log.Printf("Cannot create '%s' command: %s", c.Name, err)
			continue
		}
		log.Printf("Added command '%s' in guild '%s'", c.Name, guildID)
	}
}

// CreateGlobalCommands deletes old global commands and adds all new ones. If you
// don't want to delete you can add commands manually. s: the session you want to
// create the commands on.
func (r *Router) CreateGlobalCommands(s *discordgo.Session) {
	oldCmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Printf("error retrieving global commands: %s", err)
	}

	if len(oldCmds) > 0 {
		for _, c := range oldCmds {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", c.ID)
			if err != nil {
				log.Printf("error deleting old global command '%s': %s", c.ID, err)
			}
		}
	}

	for _, c := range r.GlobalCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.ApplicationCommand)
		if err != nil {
			log.Printf("cannot create '%s' command: %s", c.Name, err)
			continue
		}
		log.Printf("Added global command '%s'", c.Name)
	}
}

// DefaultCommandExecutor adds the router's default command executor to the session.
// This will check InteractionCreate events for your commands being called, apply middlewares,
// and run the called command's handler. session: The session you want to add the executor to.
// vars: Can be anything you want (like a database you want your commands to have access to).
func (r *Router) DefaultCommandExecutor(session *discordgo.Session, vars interface{}) {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if c, ok := r.findCommand(i.ApplicationCommandData().Name); ok {
				log.Printf("ApplicationCommand invoke: '%s'", i.ApplicationCommandData().Name)
				err := c.ApplyMiddlewares()(&Context{s, i.Interaction, vars})
				if err != nil {
					log.Printf("ApplicationCommand invoke error: %s", err)
				}
			}
		}
	})
}

// StartWithDefault adds handlers to session. Call this before openning session.
// session: The session you want to run the router on. vars: Can be anything
// you want (like a database you want your commands to have access to).
func (r *Router) StartWithDefault(session *discordgo.Session, vars interface{}) {
	r.DefaultCommandExecutor(session, vars)
	session.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) {
		r.CreateGuildCommands(s, g.ID)
	})
	session.AddHandlerOnce(func(s *discordgo.Session, e *discordgo.Ready) {
		r.CreateGlobalCommands(s)
		m := fmt.Sprintf("Logged in as %v\n", e.User)
		log.Print(m)
		for i := -19; i < len(m); i++ {
			fmt.Print("-")
		}
		fmt.Println()
	})
	log.Print("Added router handlers to session")
}

func (r *Router) findCommand(name string) (*Command, bool) {
	for _, c := range r.GlobalCommands {
		if c.Name == name {
			return &c, true
		}
	}
	for _, c := range r.GuildCommands {
		if c.Name == name {
			return &c, true
		}
	}
	return nil, false
}
