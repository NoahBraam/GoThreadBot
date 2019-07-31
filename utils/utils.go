package utils

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

// HandleThreadReaction handles a reaction of ':thread:' to a message
func HandleThreadReaction(sess *dg.Session, reaction *dg.MessageReaction) {
	// Check if thread-specific role exists

	// Check if channel exists

	// Give user correct role
}

// This function create the specific role for this thread channel
func createThreadSpecificRole(sess *dg.Session, react *dg.MessageReaction) {
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
