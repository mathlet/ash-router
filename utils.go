package ashrouter

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func WaitForMessage(s *discordgo.Session) chan *discordgo.MessageCreate {
	channel := make(chan *discordgo.MessageCreate)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		channel <- e
	})
	return channel
}

func WaitForUserMessage(s *discordgo.Session, userID string) chan *discordgo.MessageCreate {
	channel := make(chan *discordgo.MessageCreate)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		if e.Author.ID == userID {
			channel <- e
		}
	})
	return channel
}

func WaitForUserReact(s *discordgo.Session, userID string) chan *discordgo.MessageReactionAdd {
	channel := make(chan *discordgo.MessageReactionAdd)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		if e.UserID == userID {
			channel <- e
		}
	})
	return channel
}

func HasPerm(userPerms, hasPerm int64) bool {
	return userPerms&hasPerm != 0
}

func IsChannelAdmin(s *discordgo.Session, channelID, userID string) bool {
	p, err := s.UserChannelPermissions(userID, channelID)
	if err != nil {
		log.Printf("Error retrieving user channel permissions: %s", err)
		return false
	}
	if HasPerm(p, discordgo.PermissionAdministrator) {
		return true
	}
	return false
}

func CanBan(s *discordgo.Session, guildID string, m *discordgo.Member) bool {
	for _, r := range m.Roles {
		role, err := s.State.Role(guildID, r)
		if err != nil {
			log.Printf("Error retrieving '%s' role: %s", r, err)
			return false
		}
		if HasPerm(role.Permissions, discordgo.PermissionBanMembers) || HasPerm(role.Permissions, discordgo.PermissionAdministrator) {
			return true
		}
	}
	return false
}
