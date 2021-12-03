# ashrouter
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mathlet/ashrouter) [![Go Report Card](https://goreportcard.com/badge/github.com/mathlet/ashrouter)](https://goreportcard.com/report/github.com/mathlet/ashrouter)

## Getting Started

### Installation
```sh
go get github.com/mathlet/ashrouter
```

There's no set architecture but a good way to do it is to put commands in individual files. 

### Usage

Import:
```go
import "github.com/mathlet/ashrouter"
```

Create a new router:
```go
router := ashrouter.NewRouter()
```

Add global commands (call `AddGuildCommands` if you don't want the commands to be global):
```go
router.AddGlobalCommands(
		ashrouter.Command{
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
		},
	)
```
Add default router handlers to session:
```go
router.StartWithDefault()
```
This should be called before opening the session. You don't have to call this. This just makes it easy if you don't want to setup custom functionality for how the commands are loaded.


For more examples check the `examples` directory.
