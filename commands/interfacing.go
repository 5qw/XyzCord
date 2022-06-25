package commands

import "github.com/5qw/discordgo"

type ClientState interface {
	GetSelectedGuild() *discordgo.Guild
	GetSelectedChannel() *discordgo.Channel
}
