package model

import (
	"time"
)

type ReplyModel struct {
	Id int
	TechId int
	UserId int
	Content string
	Status string
	CreatedAt time.Time
}

func (ReplyModel) TableName() string {
	return "yz_tech_reply"
}