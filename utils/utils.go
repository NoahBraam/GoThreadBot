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
	roleAndChannelName := "thread-" + reaction.MessageID
	// Check if thread-specific role exists
	createThreadSpecificRole(sess, roleAndChannelName, reaction.GuildID)
	// Check if channel exists
	createThreadSpecificChannel(sess, roleAndChannelName, reaction.GuildID)
	// Give user correct role
}

// This function creates the specific role for viewing the thread
func createThreadSpecificRole(sess *dg.Session, roleName, guild string) string {
	guildRoles, _ := sess.GuildRoles(guild)
	for _, role := range guildRoles {
		if role.Name == roleName {
			return ""
		}
	}
	newRole, _ := sess.GuildRoleCreate(guild)
	newRole, _ = sess.GuildRoleEdit(guild, newRole.ID, roleName, 0, false, 0, false)
	return newRole.ID
}

// This function creates the specific channel for this thread
func createThreadSpecificChannel(sess *dg.Session, channelName, guild string) {
	channels, _ := sess.GuildChannels(guild)
	for _, channel := range channels {
		if channel.Name == channelName {
			return
		}
	}
	// TODO: permission set for @everyone so they can't read or write
	// and that specific role can read and write.
	sess.GuildChannelCreate(guild, channelName, dg.ChannelTypeGuildText)
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
