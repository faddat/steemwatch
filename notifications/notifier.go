package notifications

import (
	"io"

	"github.com/tchap/steemwatch/notifications/events"
	"github.com/tchap/steemwatch/notifications/notifiers/slack"

	"gopkg.in/mgo.v2/bson"
)

var availableNotifiers = map[string]Notifier{
	"slack": slack.NewNotifier(),
}

type Notifier interface {
	DispatchUserMentionedEvent(userId string, userSettings bson.Raw, event *events.UserMentioned) error
	DispatchTransferMadeEvent(userId string, userSettings bson.Raw, event *events.TransferMade) error
	DispatchStoryPublishedEvent(userId string, userSettings bson.Raw, event *events.StoryPublished) error
	DispatchStoryVotedEvent(userId string, userSettings bson.Raw, event *events.StoryVoted) error
	DispatchCommentPublishedEvent(userId string, userSettings bson.Raw, event *events.CommentPublished) error
	DispatchCommentVotedEvent(userId string, userSettings bson.Raw, event *events.CommentVoted) error

	io.Closer
}
