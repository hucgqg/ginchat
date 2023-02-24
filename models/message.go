package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	// 发送者
	FromId uint
	// 接收者
	TargetId uint
	// 发送类型（私聊、群聊、广播）
	Type string
	// 消息类型（文字、图片、音频）
	Media int
	// 消息内容
	Content string
	Pic     string
	Url     string
	Desc    string
	// 其它数字统计
	Amount int
}

func (table *Message) TableName() string { return "message" }
