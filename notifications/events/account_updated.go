package events

import (
	"github.com/go-steem/rpc/apis/database"
)

type AccountUpdated struct {
	Op *database.AccountUpdateOperation
}

type AccountUpdatedEventMiner struct{}

func NewAccountUpdatedEventMiner() *AccountUpdatedEventMiner {
	return &AccountUpdatedEventMiner{}
}

func (miner *AccountUpdatedEventMiner) MineEvent(
	operation *database.Operation,
	content *database.Content, // nil
) ([]interface{}, error) {

	op, ok := operation.Body.(*database.AccountUpdateOperation)
	if !ok {
		return nil, nil
	}
	return []interface{}{&AccountUpdated{op}}, nil
}
