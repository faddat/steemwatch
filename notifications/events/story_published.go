package events

import (
	"github.com/go-steem/rpc/apis/database"
)

type StoryPublished struct {
	Op      *database.CommentOperation
	Content *database.Content
}

type StoryPublishedEventMiner struct{}

func NewStoryPublishedEventMiner() *StoryPublishedEventMiner {
	return &StoryPublishedEventMiner{}
}

func (miner *StoryPublishedEventMiner) MineEvent(
	operation *database.Operation,
	content *database.Content,
) []interface{} {

	if !content.IsStory() {
		return nil
	}

	op, ok := operation.Body.(*database.CommentOperation)
	if !ok {
		return nil
	}

	return []interface{}{&StoryPublished{op, content}}
}
