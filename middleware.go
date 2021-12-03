package ashrouter

import "github.com/bwmarrin/discordgo"

// Middlewares wrap your command handlers to perform checks before the root handler is called.
type Middleware func(CommandHandler) CommandHandler

// IsAdmin is a builtin middleware that checks if the caller of the command is an admin.
// It will only run the root handler if they are. If not, it will inform them they lack permissions.
func IsAdmin(h CommandHandler) CommandHandler {
	return func(ctx *Context) error {
		if IsChannelAdmin(ctx.Session, ctx.Interaction.ChannelID, ctx.Interaction.Member.User.ID) {
			return h(ctx)
		} else {
			err := ctx.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   64,
					Content: "You must be an admin to run this command!",
				},
			})
			return err
		}
	}
}
