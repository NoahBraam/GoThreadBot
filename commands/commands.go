package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

// HelpCommand is the command to display help message
func HelpCommand(sess *dg.Session, msg *dg.Message) {
	messageStr := mentionUser(msg) + ", hi I'm ThreadBot!"
	sendMessage(sess, msg.ChannelID, messageStr)
	messageStr = "To use me, just react to a message using :thread: and I'll handle the rest!"
	sendMessage(sess, msg.ChannelID, messageStr)
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
