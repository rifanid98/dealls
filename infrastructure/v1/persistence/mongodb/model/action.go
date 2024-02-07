package model

import (
	"dealls/core/v1/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Action struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   string             `bson:"user_id"`
	TargetId string             `bson:"target_id"`
	Action   int                `bson:"action"`
	History  []ActionHistory    `bson:"history"`
	Created  time.Time          `bson:"created"`
	Modified time.Time          `bson:"modified"`
}

type ActionHistory struct {
	Action    int       `bson:"code"`
	Timestamp time.Time `bson:"timestamp"`
}

func (doc *Action) Bind(account *entity.Action) *Action {
	var history []ActionHistory
	for _, h := range account.History {
		history = append(history, ActionHistory(h))
	}

	return &Action{
		Id:       GetObjectId(account.Id),
		UserId:   account.UserId,
		TargetId: account.TargetId,
		Action:   account.Action,
		History:  history,
		Created:  account.Created,
		Modified: account.Modified,
	}
}

func (doc *Action) Entity() *entity.Action {
	var history []entity.ActionHistory
	for _, h := range doc.History {
		history = append(history, entity.ActionHistory(h))
	}

	return &entity.Action{
		Id:       GetObjectIdHex(doc.Id),
		UserId:   doc.UserId,
		TargetId: doc.TargetId,
		Action:   doc.Action,
		History:  history,
		Created:  doc.Created,
		Modified: doc.Modified,
	}
}

type Actions []Action

func (accs Actions) Bind(accounts []entity.Action) Actions {
	for i := range accounts {
		now := time.Now()
		acc := new(Action).Bind(&accounts[i])
		acc.Created = now
		acc.Modified = now
		if accounts[i].Id != "" {
			acc.Id = GetObjectId(accounts[i].Id)
		}
		accs = append(accs, *acc)
	}
	return accs
}

func (accs Actions) Generics() []any {
	var data []any
	for i := range accs {
		data = append(data, accs[i])
	}
	return data
}

func (accs Actions) Entities() []entity.Action {
	var es []entity.Action
	for i := range accs {
		es = append(es, *accs[i].Entity())
	}
	return es
}
