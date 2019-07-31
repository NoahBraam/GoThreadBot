package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/NoahBraam/GoThreadBot/utils"

	"github.com/bwmarrin/discordgo"
)

// Parameters from flag.
var accountToken string

func init() {
	// Parse command line arguments.
	flag.StringVar(&accountToken, "t", "", "Bot account token")
	flag.Parse()
	if accountToken == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	var session *discordgo.Session
	session, err = discordgo.New("Bot " + accountToken)
	setupHandlers(session)
	panicOnErr(err)
	err = session.Open()
	panicOnErr(err)
	<-make(chan struct{})
}

func setupHandlers(sess *discordgo.Session) {
	fmt.Println("Setup.")
	sess.AddHandler(messageHandler)
	sess.AddHandler(reactionHandler)
}

func messageHandler(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	message := evt.Message
	switch strings.ToLower(strings.TrimSpace(message.Content)) {
	case "~help":
		utils.HelpCommand(sess, message)
	}
}

func reactionHandler(sess *discordgo.Session, evt *discordgo.MessageReactionAdd) {
	reaction := evt.MessageReaction
	if name := reaction.Emoji.Name; name == "thread" {
		// Handle new thread channel
		utils.HandleThreadReaction(sess, reaction)
	}

}
