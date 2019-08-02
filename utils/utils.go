package utils

import (
	"fmt"

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
	roleAndChannelName := "thread-" + reaction.MessageID
	// Check if thread-specific role exists
	createThreadSpecificRole(sess, roleAndChannelName, reaction.GuildID)
	// Check if channel exists
	createThreadSpecificChannel(sess, roleAndChannelName, reaction.GuildID)
	// Give user correct role
}

// This function creates the specific role for viewing the thread
func createThreadSpecificRole(sess *dg.Session, roleName, guild string) {
	fmt.Println(roleName, guild)
}

// This function creates the specific channel for this thread
func createThreadSpecificChannel(sess *dg.Session, channelName, guild string) {
	//TODO: Check if exists first...

	sess.GuildChannelCreate(guild, channelName, dg.ChannelTypeGuildText)
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
