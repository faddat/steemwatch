package notifications

import (
	"io"
	"os"

	"github.com/tchap/steemwatch/notifications/events"
	"github.com/tchap/steemwatch/notifications/notifiers/slack"
	"github.com/tchap/steemwatch/notifications/notifiers/steemitchat"

	"gopkg.in/mgo.v2/bson"
)

var availableNotifiers = map[string]Notifier{
	"slack": slack.NewNotifier(),
}

// XXX: Ugly. Would be better to pass the values directly somehow.
func init() {
	mustGetenv := func(key string) string {
		v := os.Getenv(key)
		if v == "" {
			panic(key + " is not set")
		}
		return v
	}

	userID := mustGetenv("STEEMWATCH_STEEMIT_CHAT_USER_ID")
	authToken := mustGetenv("STEEMWATCH_STEEMIT_CHAT_AUTH_TOKEN")

	availableNotifiers["steemit-chat"] = steemitchat.NewNotifier(userID, authToken)
}

type Notifier interface {
	DispatchAccountUpdatedEvent(userId string, userSettings bson.Raw, event *events.AccountUpdated) error
	DispatchAccountWitnessVotedEvent(userId string, userSettings bson.Raw, event *events.AccountWitnessVoted) error
	DispatchTransferMadeEvent(userId string, userSettings bson.Raw, event *events.TransferMade) error
	DispatchUserMentionedEvent(userId string, userSettings bson.Raw, event *events.UserMentioned) error
	DispatchUserFollowStatusChangedEvent(userId string, userSettings bson.Raw, event *events.UserFollowStatusChanged) error
	DispatchStoryPublishedEvent(userId string, userSettings bson.Raw, event *events.StoryPublished) error
	DispatchStoryVotedEvent(userId string, userSettings bson.Raw, event *events.StoryVoted) error
	DispatchCommentPublishedEvent(userId string, userSettings bson.Raw, event *events.CommentPublished) error
	DispatchCommentVotedEvent(userId string, userSettings bson.Raw, event *events.CommentVoted) error

	io.Closer
}
