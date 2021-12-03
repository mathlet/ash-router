package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mathlet/ashrouter"
)

var Echo = ashrouter.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "echo",
		Description: "Echos text",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "text",
				Description: "The text you want to echo",
				Required:    true,
			},
		},
	},
	Handler: func(ctx *ashrouter.Context) error {
		text := ctx.ApplicationCommandData().Options[0].StringValue()
		if err := ctx.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   64,
				Content: fmt.Sprintf("Echoed `%s`", text),
			},
		}); err != nil {
			return err
		}
		_, err := ctx.ChannelMessageSend(ctx.ChannelID, text)
		return err
	},
	Middlewares: []ashrouter.Middleware{ashrouter.IsAdmin},
}
