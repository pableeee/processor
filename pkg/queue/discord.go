package queue

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordPusher struct {
	session *discordgo.Session
	channel string
}

func NewDiscordPusher(token, channelID string) (Pusher, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, fmt.Errorf("error creating session %w", err)
	}

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("error connecting %w", err)
	}

	return &DiscordPusher{session: session, channel: channelID}, nil
}

func WithSession(session *discordgo.Session, channelID string) (Pusher, error) {
	return &DiscordPusher{session: session, channel: channelID}, nil
}

func (d *DiscordPusher) Push(b []byte) error {
	_, err := d.session.ChannelMessageSend(d.channel, string(b))
	if err != nil {
		return fmt.Errorf("error sending message %w", err)
	}

	return nil
}
