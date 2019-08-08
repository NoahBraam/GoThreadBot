package utils

import (
	dg "github.com/bwmarrin/discordgo"
)

// View channel && Send messages permission
var permissions = 0x400 | 0x800

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
	roleID := createThreadSpecificRole(sess, roleAndChannelName, reaction.GuildID)
	// Check if channel exists
	createThreadSpecificChannel(sess, roleAndChannelName, reaction.GuildID, roleID)
	// Give user correct role
	addRoleToUser(sess, reaction, roleID)
}

// HandleReactionRemoved handles when the thread reaction gets removed
func HandleReactionRemoved(sess *dg.Session, reaction *dg.MessageReaction) {
	threadName := "thread-" + reaction.MessageID

	guildRoles, _ := sess.GuildRoles(reaction.GuildID)

	var roleID string
	for _, role := range guildRoles {
		if role.Name == threadName {
			roleID = role.ID
		}
	}

	member, _ := sess.GuildMember(reaction.GuildID, reaction.UserID)
	roles := member.Roles
	var newRoles []string
	for _, role := range roles {
		if roleID != role {
			newRoles = append(newRoles, role)
		}
	}

	sess.GuildMemberEdit(reaction.GuildID, reaction.UserID, newRoles)
}

func addRoleToUser(sess *dg.Session, react *dg.MessageReaction, roleID string) {
	member, _ := sess.GuildMember(react.GuildID, react.UserID)
	roles := member.Roles
	roles = append(roles, roleID)

	sess.GuildMemberEdit(react.GuildID, react.UserID, roles)
}

// This function creates the specific role for viewing the thread
func createThreadSpecificRole(sess *dg.Session, roleName, guild string) string {
	guildRoles, _ := sess.GuildRoles(guild)
	for _, role := range guildRoles {
		if role.Name == roleName {
			return role.ID
		}
	}
	newRole, _ := sess.GuildRoleCreate(guild)
	newRole, _ = sess.GuildRoleEdit(guild, newRole.ID, roleName, 0, false, permissions, false)
	return newRole.ID
}

// This function creates the specific channel for this thread
func createThreadSpecificChannel(sess *dg.Session, channelName, guild, roleID string) {
	channels, _ := sess.GuildChannels(guild)
	for _, channel := range channels {
		if channel.Name == channelName {
			return
		}
	}
	// TODO: permission set for @everyone so they can't read or write
	// and that specific role can read and write.
	newChannel, _ := sess.GuildChannelCreate(guild, channelName, dg.ChannelTypeGuildText)
	everyonePermission := &dg.PermissionOverwrite{
		ID:   guild,
		Type: "role",
		Deny: permissions}
	rolePermission := &dg.PermissionOverwrite{
		ID:    roleID,
		Type:  "role",
		Allow: permissions}

	channelEdit := &dg.ChannelEdit{}
	channelEdit.PermissionOverwrites = append(channelEdit.PermissionOverwrites, everyonePermission)
	channelEdit.PermissionOverwrites = append(channelEdit.PermissionOverwrites, rolePermission)
	sess.ChannelEditComplex(newChannel.ID, channelEdit)
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
