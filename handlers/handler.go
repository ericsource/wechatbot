package handlers

import (
	"github.com/eatmoreapple/openwechat"
	"log"
	"time"
)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

type HandlerType string

const (
	GroupHandler = "group"
	UserHandler  = "user"
)

// handlers 所有消息类型类型的处理器
var handlers map[HandlerType]MessageHandlerInterface

func init() {
	handlers = make(map[HandlerType]MessageHandlerInterface)
	handlers[GroupHandler] = NewGroupMessageHandler()
	handlers[UserHandler] = NewUserMessageHandler()
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	//log.Printf("hadler Received msg : %v", msg.Content)

	// 好友申请
	if msg.IsFriendAdd() {
		//if config.LoadConfig().AutoPass {
		//	_, err := msg.Agree("你好我是基于chatGPT引擎开发的微信机器人，你可以向我提问任何问题。")
		//	if err != nil {
		//		log.Fatalf("add friend agree error : %v", err)
		//		return
		//	}
		//}
		return
	}

	// 判断是否自己发送
	if msg.IsSendBySelf() {
		return
	}

	// 空消息处理
	if msg.Content == "" {
		log.Println("msg content empty")
		return
	}

	//fmt.Println("create time:", msg.CreateTime, msg., time.Now().Unix())
	// 旧内容不要处理
	if (time.Now().Unix() - msg.CreateTime > 5) {
		return
	}

	// 处理群消息
	if msg.IsSendByGroup() {
		handlers[GroupHandler].handle(msg)
		return
	}

	// 私聊文本
	if msg.IsText() {
		handlers[UserHandler].handle(msg)
		return
	}

	// 私聊图片 todo
	if msg.IsPicture() {
		//handlers[UserHandler].handle(msg)
		return
	}

	return
}
