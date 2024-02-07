package entity

import (
	"time"
)

type Action struct {
	Id       string          `json:"_id"`
	UserId   string          `json:"user_id"`
	TargetId string          `json:"target_id"`
	Action   int             `json:"action"`
	History  []ActionHistory `json:"history"`
	Created  time.Time       `json:"created"`
	Modified time.Time       `json:"modified"`
}

type ActionHistory struct {
	Action    int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}
