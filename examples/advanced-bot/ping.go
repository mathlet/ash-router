package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mathlet/ashrouter"
)

var Ping = ashrouter.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Pings the bot",
		Type:        discordgo.ChatApplicationCommand,
	},
	Handler: func(ctx *ashrouter.Context) error {
		err := ctx.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   64,
				Content: fmt.Sprintf("Pong %dms!", ctx.HeartbeatLatency()/time.Millisecond),
			},
		})
		return err
	},
}
